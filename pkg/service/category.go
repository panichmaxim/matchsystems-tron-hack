package service

import (
	"context"
	"gitlab.com/rubin-dev/api/pkg/models"
	"gitlab.com/rubin-dev/api/pkg/store"
)

var _ CategoryService = (*serviceImpl)(nil)

type CategoryService interface {
	store.CategoryStore
}

func (s *serviceImpl) CategoryList(ctx context.Context) ([]*models.Category, int, error) {
	return s.s.CategoryList(ctx)
}

func (s *serviceImpl) CategoryFindByID(ctx context.Context, id int64) (*models.Category, error) {
	return s.s.CategoryFindByID(ctx, id)
}

func (s *serviceImpl) CategoryCreate(ctx context.Context, category *models.Category) error {
	return s.s.CategoryCreate(ctx, category)
}

func (s *serviceImpl) CategoryUpdate(ctx context.Context, category *models.Category, columns ...string) error {
	return s.s.CategoryUpdate(ctx, category, columns...)
}

func (s *serviceImpl) CategoryRemoveByID(ctx context.Context, id int64) error {
	return s.s.CategoryRemoveByID(ctx, id)
}
