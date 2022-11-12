package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"gitlab.com/rubin-dev/api/internal/graph/generated"
	"gitlab.com/rubin-dev/api/internal/graph/model"
	"gitlab.com/rubin-dev/api/internal/tools"
	"gitlab.com/rubin-dev/api/pkg/models"
	"gitlab.com/rubin-dev/api/pkg/validator"
)

func (r *billingPacketResolver) CreatedAt(ctx context.Context, obj *models.BillingPacket) (string, error) {
	return obj.CreatedAt.Format(time.RFC3339), nil
}

func (r *billingRequestResolver) Categories(ctx context.Context, obj *models.BillingRequest) ([]*models.BillingRisk, error) {
	return r.svc.BillingRisks(ctx, obj.ID)
}

func (r *billingRequestResolver) CreatedAt(ctx context.Context, obj *models.BillingRequest) (string, error) {
	return obj.CreatedAt.Format(time.RFC3339), nil
}

func (r *billingRiskResolver) Directory(ctx context.Context, obj *models.BillingRisk) (*models.CategoryGroup, error) {
	return models.FindDirectory(obj.DirectoryID), nil
}

func (r *billingRiskResolver) Category(ctx context.Context, obj *models.BillingRisk) (*models.Category, error) {
	if obj.CategoryID == nil {
		return nil, nil
	}

	return r.svc.CategoryFindByNumber(ctx, *obj.CategoryID)
}

func (r *billingStatisticsBlockchainResolver) Categories(ctx context.Context, obj *models.BillingStatisticsBlockchain) ([]*models.BillingStatisticsCategory, error) {
	var cats []*models.BillingStatisticsCategory
	for n, v := range obj.Categories {
		cats = append(cats, &models.BillingStatisticsCategory{
			Name: n,
			Risk: float64(v),
		})
	}
	return cats, nil
}

func (r *billingStatisticsCategoryResolver) Number(ctx context.Context, obj *models.BillingStatisticsCategory) (int, error) {
	panic(fmt.Errorf("not implemented"))
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

func (r *queryResolver) BillingHistoryList(ctx context.Context, filter model.BillingHistoryListInput) (*model.BillingHistoryListResponse, error) {
	u := WithUser(ctx)

	fromTime, toTime := tools.ParseFromToDates(filter.From, filter.To)

	resp, count, err := r.svc.BillingHistoryList(ctx, u.ID, filter.Page, filter.PageSize, fromTime, toTime, filter.Last)
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

func (r *queryResolver) BillingStatistics(ctx context.Context, filter model.BillingStatisticsFilterInput) (*model.BillingStatisticsResponse, error) {
	u := WithUser(ctx)

	fromTime, toTime := tools.ParseFromToDates(filter.From, filter.To)
	stats, err := r.svc.BillingStatisticsNetwork(ctx, u.ID, fromTime, toTime, filter.Network, filter.Last)
	if err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.BillingStatisticsResponse{Errors: errs}, nil
		}

		return nil, err
	}

	return &model.BillingStatisticsResponse{Stats: stats}, nil
}

func (r *queryResolver) BillingStatisticsSummary(ctx context.Context, filter model.StatisticsSummaryInput) (*model.BillingStatisticsSummaryResponse, error) {
	u := WithUser(ctx)
	fromTime, toTime := tools.ParseFromToDates(filter.From, filter.To)
	items, err := r.svc.BillingStatisticsSummary(ctx, u.ID, fromTime, toTime, filter.Last)
	if err != nil {
		return nil, err
	}

	return &model.BillingStatisticsSummaryResponse{
		Items: items,
	}, nil
}

func (r *queryResolver) BillingStatisticsRiskRange(ctx context.Context, filter model.BillingStatisticsRiskRangeInput) ([]int, error) {
	u := WithUser(ctx)
	fromTime, toTime := tools.ParseFromToDates(filter.From, filter.To)
	return r.svc.BillingStatisticsRiskRanges(
		ctx,
		u.ID,
		fromTime,
		toTime,
		filter.Network,
		filter.Last,
	)
}

// BillingPacket returns generated.BillingPacketResolver implementation.
func (r *Resolver) BillingPacket() generated.BillingPacketResolver { return &billingPacketResolver{r} }

// BillingRequest returns generated.BillingRequestResolver implementation.
func (r *Resolver) BillingRequest() generated.BillingRequestResolver {
	return &billingRequestResolver{r}
}

// BillingRisk returns generated.BillingRiskResolver implementation.
func (r *Resolver) BillingRisk() generated.BillingRiskResolver { return &billingRiskResolver{r} }

// BillingStatisticsBlockchain returns generated.BillingStatisticsBlockchainResolver implementation.
func (r *Resolver) BillingStatisticsBlockchain() generated.BillingStatisticsBlockchainResolver {
	return &billingStatisticsBlockchainResolver{r}
}

// BillingStatisticsCategory returns generated.BillingStatisticsCategoryResolver implementation.
func (r *Resolver) BillingStatisticsCategory() generated.BillingStatisticsCategoryResolver {
	return &billingStatisticsCategoryResolver{r}
}

type billingPacketResolver struct{ *Resolver }
type billingRequestResolver struct{ *Resolver }
type billingRiskResolver struct{ *Resolver }
type billingStatisticsBlockchainResolver struct{ *Resolver }
type billingStatisticsCategoryResolver struct{ *Resolver }
