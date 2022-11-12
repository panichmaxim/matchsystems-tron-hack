package store

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"gitlab.com/rubin-dev/api/pkg/models"
	"strings"
)

var _ CategoryStore = (*storeImpl)(nil)

type CategoryStore interface {
	CategoryFindByID(ctx context.Context, id int64) (*models.Category, error)
	CategoryFindByNumber(ctx context.Context, id int) (*models.Category, error)
	CategoryCreate(ctx context.Context, category *models.Category) error
	CategoryUpdate(ctx context.Context, category *models.Category, columns ...string) error
	CategoryRemoveByID(ctx context.Context, id int64) error
	CategoryList(ctx context.Context, id *int64) (u []*models.Category, err error)
	CategoryAllList(ctx context.Context) (u []*models.Category, err error)
	CategoryGroupList() []*models.CategoryGroup
	CategoryFindByName(ctx context.Context, name string) (*models.Category, error)
}

func (s *storeImpl) CategoryGroupList() []*models.CategoryGroup {
	return models.GetCategoryGroups()
}

func (s *storeImpl) CategoryList(ctx context.Context, id *int64) (u []*models.Category, err error) {
	q := s.db.
		NewSelect().
		Model((*models.Category)(nil))
	if id != nil {
		q = q.Where("category_group_id = ?", id)
	} else {
		q = q.Where("category_group_id IS NULL")
	}
	err = q.Order("id").Scan(ctx, &u)
	return
}

func (s *storeImpl) CategoryAllList(ctx context.Context) (u []*models.Category, err error) {
	err = s.db.
		NewSelect().
		Model((*models.Category)(nil)).
		Order("id").Scan(ctx, &u)
	return
}

func (s *storeImpl) CategoryFindByName(ctx context.Context, name string) (*models.Category, error) {
	f := new(models.Category)
	err := s.db.
		NewSelect().
		Model(f).
		Where("LOWER(name) = ?", strings.ToLower(name)).
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return f, nil
}

func (s *storeImpl) CategoryFindByNumber(ctx context.Context, number int) (*models.Category, error) {
	f := new(models.Category)
	err := s.db.
		NewSelect().
		Model(f).
		Where("number = ?", number).
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return f, nil
}

func (s *storeImpl) CategoryFindByID(ctx context.Context, id int64) (*models.Category, error) {
	f := new(models.Category)
	err := s.db.
		NewSelect().
		Model(f).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return f, nil
}

func (s *storeImpl) CategoryUpdate(ctx context.Context, category *models.Category, columns ...string) error {
	_, err := s.db.
		NewUpdate().
		Model(category).
		WherePK().
		Column(columns...).
		Returning("*").
		Exec(ctx)
	return err
}

func (s *storeImpl) CategoryCreate(ctx context.Context, category *models.Category) error {
	_, err := s.db.
		NewInsert().
		Model(category).
		Exec(ctx)
	return err
}

func (s *storeImpl) CategoryRemoveByID(ctx context.Context, id int64) error {
	_, err := s.db.
		NewDelete().
		Model((*models.Category)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
