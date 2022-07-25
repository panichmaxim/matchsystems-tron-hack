package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"gitlab.com/rubin-dev/api/internal/graph/generated"
)

func (r *mutationResolver) Health(ctx context.Context) (bool, error) {
	if err := r.svc.Health(ctx); err != nil {
		return false, err
	}

	return true, nil
}

func (r *queryResolver) Health(ctx context.Context) (bool, error) {
	if err := r.svc.Health(ctx); err != nil {
		return false, err
	}

	return true, nil
}

func (r *subscriptionResolver) Health(ctx context.Context) (<-chan bool, error) {
	if err := r.svc.Health(ctx); err != nil {
		return nil, err
	}

	return nil, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
