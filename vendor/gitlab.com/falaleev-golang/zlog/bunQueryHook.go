package zlog

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/uptrace/bun"
)

var _ bun.QueryHook = (*QueryHook)(nil)

type QueryHook struct{}

// BeforeQuery before query zerolog hook.
func (h *QueryHook) BeforeQuery(ctx context.Context, _ *bun.QueryEvent) context.Context {
	return ctx
}

// AfterQuery after query zerolog hook.
func (h *QueryHook) AfterQuery(ctx context.Context, event *bun.QueryEvent) {
	var logEvent *zerolog.Event

	// do not log sql.ErrNoRows as real error
	l := zerolog.Ctx(ctx)
	if errors.Is(event.Err, sql.ErrNoRows) {
		logEvent = l.Warn().Err(event.Err)
	} else {
		logEvent = l.Err(event.Err)
	}

	logEvent.
		Str("type", queryOperation(event.Query)).
		Str("query", event.Query).
		Str("operation", eventOperation(event)).
		Str("duration", time.Since(event.StartTime).String()).
		Msg("query")
}

func eventOperation(event *bun.QueryEvent) string {
	switch event.IQuery.(type) {
	case *bun.SelectQuery:
		return "SELECT"
	case *bun.InsertQuery:
		return "INSERT"
	case *bun.UpdateQuery:
		return "UPDATE"
	case *bun.DeleteQuery:
		return "DELETE"
	case *bun.CreateTableQuery:
		return "CREATE TABLE"
	case *bun.DropTableQuery:
		return "DROP TABLE"
	default:
		return queryOperation(event.Query)
	}
}

const maxOperationNameSize = 16

func queryOperation(name string) string {
	if idx := strings.Index(name, " "); idx > 0 {
		name = name[:idx]
	}

	if len(name) > maxOperationNameSize {
		name = name[:maxOperationNameSize]
	}

	return name
}
