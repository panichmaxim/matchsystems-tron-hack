package store

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"gitlab.com/rubin-dev/api/pkg/models"
)

var _ CategoryStore = (*storeImpl)(nil)

type CategoryStore interface {
	CategoryFindByID(ctx context.Context, id int64) (*models.Category, error)
	CategoryCreate(ctx context.Context, category *models.Category) error
	CategoryUpdate(ctx context.Context, category *models.Category, columns... string) error
	CategoryRemoveByID(ctx context.Context, id int64) error
	CategoryList(ctx context.Context) (u []*models.Category, count int, err error)
}

func (s *storeImpl) CategoryList(ctx context.Context) (u []*models.Category, count int, err error) {
	count, err = s.db.
		NewSelect().
		Model((*models.Category)(nil)).
		Order("id").
		ScanAndCount(ctx, &u)
	return
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

func (s *storeImpl) CategoryUpdate(ctx context.Context, category *models.Category, columns... string) error {
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
