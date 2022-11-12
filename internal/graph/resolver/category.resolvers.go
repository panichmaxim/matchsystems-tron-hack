package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"gitlab.com/rubin-dev/api/internal/graph/generated"
	"gitlab.com/rubin-dev/api/internal/graph/model"
	"gitlab.com/rubin-dev/api/pkg/models"
	"gitlab.com/rubin-dev/api/pkg/validator"
)

func (r *categoryResolver) CategoryGroup(ctx context.Context, obj *models.Category) (*models.CategoryGroup, error) {
	if obj.CategoryGroupID == nil {
		return nil, nil
	}
	return models.FindDirectory(*obj.CategoryGroupID), nil
}

func (r *mutationResolver) CategoryCreate(ctx context.Context, input model.CategoryCreateInput) (*model.CategoryCreateResponse, error) {
	category := &models.Category{
		Name:            input.Name,
		Risk:            input.Risk,
		Number:          input.Number,
		DescriptionRu:   input.DescriptionRu,
		DescriptionEn:   input.DescriptionEn,
		CategoryGroupID: input.CategoryGroupID,
	}
	if err := r.svc.CategoryCreate(ctx, category); err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.CategoryCreateResponse{Errors: errs}, nil
		}

		return nil, err
	}

	return &model.CategoryCreateResponse{Category: category}, nil
}

func (r *mutationResolver) CategoryUpdate(ctx context.Context, id int64, input model.CategoryUpdateInput) (*model.CategoryUpdateResponse, error) {
	category, err := r.svc.CategoryFindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	var columns []string
	if input.Name != nil {
		category.Name = *input.Name
		columns = append(columns, "name")
	}
	if input.Risk != nil {
		category.Risk = *input.Risk
		columns = append(columns, "risk")
	}
	if input.DescriptionRu != nil {
		category.DescriptionRu = *input.DescriptionRu
		columns = append(columns, "description_ru")
	}
	if input.DescriptionEn != nil {
		category.DescriptionEn = *input.DescriptionEn
		columns = append(columns, "description_en")
	}
	if input.Number != nil {
		category.Number = *input.Number
		columns = append(columns, "number")
	}
	if input.CategoryGroupID != nil {
		category.CategoryGroupID = input.CategoryGroupID
		columns = append(columns, "category_group_id")
	}

	if len(columns) == 0 {
		return &model.CategoryUpdateResponse{Category: category}, nil
	}

	if err := r.svc.CategoryUpdate(ctx, category, columns...); err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.CategoryUpdateResponse{Errors: errs}, nil
		}

		return nil, err
	}

	return &model.CategoryUpdateResponse{Category: category}, nil
}

func (r *mutationResolver) CategoryRemoveByID(ctx context.Context, id int64) (*model.CategoryRemoveResponse, error) {
	if err := r.svc.CategoryRemoveByID(ctx, id); err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.CategoryRemoveResponse{Errors: errs}, nil
		}

		return nil, err
	}

	return &model.CategoryRemoveResponse{}, nil
}

func (r *queryResolver) CategoryList(ctx context.Context, id *int64) (*model.CategoryListResponse, error) {
	items, err := r.svc.CategoryList(ctx, id)
	if err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.CategoryListResponse{Errors: errs}, nil
		}

		return nil, err
	}

	return &model.CategoryListResponse{
		Edge: items,
	}, nil
}

func (r *queryResolver) CategoryAllList(ctx context.Context) (*model.CategoryListResponse, error) {
	items, err := r.svc.CategoryAllList(ctx)
	if err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.CategoryListResponse{Errors: errs}, nil
		}

		return nil, err
	}

	return &model.CategoryListResponse{
		Edge: items,
	}, nil
}

func (r *queryResolver) CategoryFindByID(ctx context.Context, id int64) (*models.Category, error) {
	return r.svc.CategoryFindByID(ctx, id)
}

// Category returns generated.CategoryResolver implementation.
func (r *Resolver) Category() generated.CategoryResolver { return &categoryResolver{r} }

type categoryResolver struct{ *Resolver }
