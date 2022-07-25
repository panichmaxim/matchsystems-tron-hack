package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"gitlab.com/rubin-dev/api/internal/graph/generated"
	"gitlab.com/rubin-dev/api/pkg/neoutils"
)

func (r *nodeResolver) Props(ctx context.Context, obj *neoutils.Node) (interface{}, error) {
	return obj.Props, nil
}

// Node returns generated.NodeResolver implementation.
func (r *Resolver) Node() generated.NodeResolver { return &nodeResolver{r} }

type nodeResolver struct{ *Resolver }
