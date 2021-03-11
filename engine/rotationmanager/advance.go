package rotationmanager

import (
	"context"
	"fmt"
	"time"

	"github.com/target/goalert/schedule/rotation"
	"github.com/target/goalert/util/log"
)

type advance struct {
	id string
	t  time.Time
	p  int
}

type rotState struct {
	rotation.State
	Version int
}

// calcAdvance will calculate rotation advancement if it is required. If not, nil is returned
func calcAdvance(ctx context.Context, t time.Time, rot *rotation.Rotation, state rotState, partCount int) *advance {
	var mustUpdate bool

	// get next shift start time
	newStart := rot.EndTime(state.ShiftStart)
	if state.Version == 1 {
		newStart = calcOldEndTime(rot, state.ShiftStart)
		mustUpdate = true
	}

	if state.Position >= partCount {
		// deleted last participant
		state.Position = 0
		mustUpdate = true
	}

	if newStart.After(t) || state.Version == 1 {
		if mustUpdate {
			return &advance{
				id: rot.ID,
				t:  state.ShiftStart,
				p:  state.Position,
			}
		}
		// in the future, so nothing to do yet
		return nil
	}

	if !newStart.After(t.Add(-15 * time.Minute)) {
		log.Log(log.WithField(ctx, "RotationID", rot.ID), fmt.Errorf("rotation advanced late (%s)", t.Sub(newStart).String()))
	}

	state.ShiftStart = newStart

	c := 0
	for {
		c++
		if c > 10000 {
			panic("too many rotation advances")
		}

		state.Position = (state.Position + 1) % partCount
		end := rot.EndTime(state.ShiftStart)
		if end.After(t) {
			break
		}
		state.ShiftStart = end
	}

	return &advance{
		id: rot.ID,
		t:  state.ShiftStart,
		p:  state.Position,
	}
}
