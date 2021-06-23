package user

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/binary"
	"strings"
	"time"

	"github.com/golang/groupcache"
	"github.com/google/uuid"
	"github.com/target/goalert/permission"
)

func timeKey(key string, dur time.Duration) string {
	return key + "\n" + time.Now().Round(dur).Format(time.RFC3339)
}
func keyName(keyWithTime string) string {
	return strings.SplitN(keyWithTime, "\n", 2)[0]
}

// ExistanceChecker allows checking if various users exist.
type ExistanceChecker interface {
	UserExistsString(id string) bool
	UserExistsUUID(id uuid.UUID) bool
}

type checker map[uuid.UUID]struct{}

func (c checker) UserExistsString(idStr string) bool {
	if idStr == "" {
		return false
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		return false
	}

	return c.UserExistsUUID(id)
}
func (c checker) UserExistsUUID(id uuid.UUID) bool { _, ok := c[id]; return ok }

// UserExists returns an ExistanceChecker.
func (db *DB) UserExists(ctx context.Context) (ExistanceChecker, error) {
	err := permission.LimitCheckAny(ctx)
	if err != nil {
		return nil, err
	}
	m, err := db.userExistMap(ctx)
	if err != nil {
		return nil, err
	}
	return checker(m), nil
}

func (db *DB) userExistMap(ctx context.Context) (map[uuid.UUID]struct{}, error) {
	var data []byte
	err := db.grp.Get(ctx, timeKey("userIDs", time.Minute), groupcache.AllocatingByteSliceSink(&data))
	if err != nil {
		return nil, err
	}

	m := <-db.userExist
	if bytes.Equal(data[:sha256.Size], db.userExistHash) {
		db.userExist <- m
		return m, nil
	}

	ids := make([]uuid.UUID, (len(data)-sha256.Size)/16)
	err = binary.Read(bytes.NewReader(data[sha256.Size:]), binary.BigEndian, &ids)
	if err != nil {
		db.userExist <- m
		return nil, err
	}

	m = make(map[uuid.UUID]struct{}, len(ids))
	for _, id := range ids {
		m[id] = struct{}{}
	}
	db.userExistHash = data[:sha256.Size]
	db.userExist <- m

	return m, nil
}

func (db *DB) currentUserIDs(ctx context.Context) (result []byte, err error) {
	rows, err := db.ids.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []uuid.UUID
	sum := sha256.New()
	for rows.Next() {
		var id uuid.UUID
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		sum.Write(id[:])
		data = append(data, id)
	}

	buf := bytes.NewBuffer(nil)
	buf.Write(sum.Sum(nil))
	err = binary.Write(buf, binary.BigEndian, data)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (db *DB) cacheGet(ctx context.Context, key string, dest groupcache.Sink) error {
	switch keyName(key) {
	case "userIDs":
		data, err := db.currentUserIDs(ctx)
		if err != nil {
			return err
		}
		return dest.SetBytes(data)
	}
	return nil
}
