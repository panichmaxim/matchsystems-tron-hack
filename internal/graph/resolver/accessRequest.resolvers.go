package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
)

func (r *mutationResolver) AccessRequest(ctx context.Context) (bool, error) {
	err := r.svc.AccessRequest(ctx, WithUser(ctx))
	return err == nil, err
}
