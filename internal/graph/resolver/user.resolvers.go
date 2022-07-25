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

func (r *queryResolver) User(ctx context.Context, id int64) (*models.User, error) {
	return r.svc.UserFindByID(ctx, id)
}

func (r *queryResolver) UserList(ctx context.Context, page int, pageSize int) (*model.UserListResponse, error) {
	resp, count, err := r.svc.UserList(ctx, &models.UserListRequest{Page: page, PageSize: pageSize})
	if err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.UserListResponse{Errors: errs}, nil
		}

		return nil, err
	}

	return &model.UserListResponse{
		Edge:  resp,
		Total: &count,
	}, nil
}

func (r *userResolver) CreatedAt(ctx context.Context, obj *models.User) (string, error) {
	return obj.CreatedAt.Format(time.RFC3339), nil
}

func (r *userResolver) UpdatedAt(ctx context.Context, obj *models.User) (string, error) {
	return obj.UpdatedAt.Format(time.RFC3339), nil
}

func (r *userResolver) Billing(ctx context.Context, obj *models.User) (*model.Billing, error) {
	billing, err := r.svc.BillingGetOrCreate(ctx, obj.ID)
	if err != nil {
		return nil, err
	}

	historyRequests, err := r.svc.BillingHistoryRequestsCount(ctx, obj.ID)
	if err != nil {
		return nil, err
	}

	return &model.Billing{
		Requests:        billing.Requests,
		HistoryRequests: historyRequests,
	}, nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
