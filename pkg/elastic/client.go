package elastic

import (
	"context"
)

type Client interface {
	Search(ctx context.Context, value string, page, limit int, wildcard bool) ([]*Entity, int, error)
	SearchCount(ctx context.Context, value string, wildcard bool) (int, error)
}
