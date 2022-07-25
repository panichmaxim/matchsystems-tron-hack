package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"gitlab.com/rubin-dev/api/internal/graph/model"
	"gitlab.com/rubin-dev/api/pkg/validator"
)

func (r *queryResolver) Search(ctx context.Context, query string, page int, limit int) (*model.SearchResponse, error) {
	items, total, err := r.svc.Search(ctx, query, page, limit)
	if err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.SearchResponse{Errors: errs}, nil
		}
		return nil, err
	}

	return &model.SearchResponse{Edge: items, Total: &total}, nil
}

func (r *queryResolver) SearchCount(ctx context.Context, query string) (*model.SearchCountResponse, error) {
	items, err := r.svc.SearchCount(ctx, query)
	if err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.SearchCountResponse{Errors: errs}, nil
		}
		return nil, err
	}

	return &model.SearchCountResponse{Count: &items}, nil
}
