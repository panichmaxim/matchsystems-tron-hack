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

func (r *billingPacketResolver) CreatedAt(ctx context.Context, obj *models.BillingPacket) (string, error) {
	return obj.CreatedAt.Format(time.RFC3339), nil
}

func (r *billingRequestResolver) CreatedAt(ctx context.Context, obj *models.BillingRequest) (string, error) {
	return obj.CreatedAt.Format(time.RFC3339), nil
}

func (r *mutationResolver) BillingAddPacket(ctx context.Context, userID int64, requests int) (*model.BillingAddPacketResponse, error) {
	user, err := r.svc.UserFindByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	billing, err := r.svc.BillingAddRequests(ctx, user.ID, requests)
	if err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.BillingAddPacketResponse{Errors: errs}, nil
		}

		return nil, err
	}

	return &model.BillingAddPacketResponse{Billing: billing}, nil
}

func (r *queryResolver) BillingHistoryList(ctx context.Context, page int, pageSize int) (*model.BillingHistoryListResponse, error) {
	u := WithUser(ctx)
	resp, count, err := r.svc.BillingHistoryList(ctx, u.ID, page, pageSize)
	if err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.BillingHistoryListResponse{Errors: errs}, nil
		}

		return nil, err
	}

	return &model.BillingHistoryListResponse{
		Edge:  resp,
		Total: &count,
	}, nil
}

// BillingPacket returns generated.BillingPacketResolver implementation.
func (r *Resolver) BillingPacket() generated.BillingPacketResolver { return &billingPacketResolver{r} }

// BillingRequest returns generated.BillingRequestResolver implementation.
func (r *Resolver) BillingRequest() generated.BillingRequestResolver {
	return &billingRequestResolver{r}
}

type billingPacketResolver struct{ *Resolver }
type billingRequestResolver struct{ *Resolver }
