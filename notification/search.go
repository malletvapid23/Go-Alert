package notification

import (
	"context"
	"database/sql"
	"text/template"
	"time"

	"github.com/pkg/errors"
	"github.com/target/goalert/permission"
	"github.com/target/goalert/search"
	"github.com/target/goalert/util/sqlutil"
	"github.com/target/goalert/validation/validate"
)

// SearchOptions allow filtering and paginating the list of messages.
type SearchOptions struct {
	Search string       `json:"s,omitempty"`
	After  SearchCursor `json:"a,omitempty"`

	// Omit specifies a list of message IDs to exclude from the results
	Omit []string `json:"o,omitempty"`

	Limit int `json:"-"`
}

// SearchCursor is used to indicate a position in a paginated list.
type SearchCursor struct {
	ID        string    `json:"i,omitempty"`
	CreatedAt time.Time `json:"n,omitempty"`
}

var searchTemplate = template.Must(template.New("search").Funcs(search.Helpers()).Parse(`
	SELECT
		om.id, om.created_at, om.last_status_at, om.message_type, om.last_status, om.status_details,
		om.src_value, om.alert_id, om.provider_msg_id,
		om.user_id, om.contact_method_id, om.channel_id, om.service_id
	FROM outgoing_messages om
	WHERE true
	{{if .Omit}}
		AND NOT om.id = any(:omit)
	{{end}}
	{{if .Search}}
		AND {{prefixSearch "search" "om.src_value"}}
	{{end}}
	{{if .After.CreatedAt}}
		AND om.created_at > :afterCreatedAt
	{{end}}
		AND om.created_at < now()
		AND om.last_status != 'bundled'
	ORDER BY om.created_at, om.id
	LIMIT {{.Limit}}
`))

type renderData SearchOptions

func (opts renderData) Normalize() (*renderData, error) {
	if opts.Limit == 0 {
		opts.Limit = 50
	}

	err := validate.Many(
		validate.Search("Search", opts.Search),
		// should be 1 more than the expected limit
		validate.Range("Limit", opts.Limit, 0, 101),
		validate.ManyUUID("Omit", opts.Omit, 50),
	)
	if err != nil {
		return nil, err
	}

	return &opts, err
}

func (opts renderData) QueryArgs() []sql.NamedArg {
	return []sql.NamedArg{
		sql.Named("search", opts.Search),
		sql.Named("afterCreatedAt", opts.After.CreatedAt),
		sql.Named("omit", sqlutil.UUIDArray(opts.Omit)),
	}
}

func (s *Store) Search(ctx context.Context, opts *SearchOptions) ([]MessageLog, error) {
	if opts == nil {
		opts = &SearchOptions{}
	}

	err := permission.LimitCheckAny(ctx, permission.System, permission.User)
	if err != nil {
		return nil, err
	}

	data, err := (*renderData)(opts).Normalize()
	if err != nil {
		return nil, err
	}

	query, args, err := search.RenderQuery(ctx, searchTemplate, data)
	if err != nil {
		return nil, errors.Wrap(err, "render query")
	}

	rows, err := s.db.QueryContext(ctx, query, args...)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []MessageLog
	var l MessageLog
	var alertID sql.NullInt64
	var chanID sqlutil.NullUUID
	var serviceID sql.NullString
	var srcValue sql.NullString
	var userID sql.NullString
	var cmID sql.NullString
	var providerID sql.NullString

	for rows.Next() {
		err = rows.Scan(
			&l.ID,
			&l.CreatedAt,
			&l.LastStatusAt,
			&l.MessageType,
			&l.LastStatus,
			&l.StatusDetails,
			&srcValue,
			&alertID,
			&providerID,
			&userID,
			&cmID,
			&chanID,
			&serviceID,
		)
		if err != nil {
			return nil, err
		}

		// set all the nullable fields
		if providerID.String != "" {
			pm, err := ParseProviderMessageID(providerID.String)
			if err != nil {
				return nil, err
			}
			l.ProviderMsgID = &pm
		}
		l.AlertID = int(alertID.Int64)
		l.ChannelID = chanID.UUID
		l.ServiceID = serviceID.String
		l.SrcValue = srcValue.String
		l.UserID = userID.String
		l.ContactMethodID = cmID.String

		result = append(result, l)
	}

	return result, nil
}
