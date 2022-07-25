package elastic

import (
	"context"
)

type Client interface {
	Search(ctx context.Context, value string, page, limit int) ([]*Entity, int, error)
	SearchCount(ctx context.Context, value string) (int, error)
}
