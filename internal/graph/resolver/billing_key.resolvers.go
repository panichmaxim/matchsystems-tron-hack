package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"gitlab.com/rubin-dev/api/internal/graph/generated"
	"gitlab.com/rubin-dev/api/internal/graph/model"
	"gitlab.com/rubin-dev/api/pkg/models"
	"gitlab.com/rubin-dev/api/pkg/validator"
)

func (r *billingKeyResolver) CreatedAt(ctx context.Context, obj *models.BillingKey) (string, error) {
	return obj.CreatedAt.Format(time.RFC3339), nil
}

func (r *mutationResolver) BillingKeyCreate(ctx context.Context) (*models.BillingKey, error) {
	u := WithUser(ctx)
	return r.svc.BillingCreateApiKey(ctx, u.ID)
}

func (r *mutationResolver) BillingKeyRemove(ctx context.Context, id int64) (bool, error) {
	u := WithUser(ctx)
	err := r.svc.BillingKeyRemove(ctx, u.ID, id)
	return err == nil, err
}

func (r *queryResolver) BillingKeyList(ctx context.Context, page int, pageSize int) (*model.BillingKeyResponse, error) {
	u := WithUser(ctx)
	resp, count, err := r.svc.BillingKeyList(ctx, u.ID, page, pageSize)
	if err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.BillingKeyResponse{Errors: errs}, nil
		}

		return nil, err
	}

	return &model.BillingKeyResponse{
		Edge:  resp,
		Total: &count,
	}, nil
}

// BillingKey returns generated.BillingKeyResolver implementation.
func (r *Resolver) BillingKey() generated.BillingKeyResolver { return &billingKeyResolver{r} }

type billingKeyResolver struct{ *Resolver }
