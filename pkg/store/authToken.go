package store

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"gitlab.com/rubin-dev/api/pkg/models"
)

var _ AuthTokenStore = (*storeImpl)(nil)

type AuthTokenStore interface {
	AuthTokenFind(ctx context.Context, id string) (*models.AuthToken, error)
	AuthTokenCreate(ctx context.Context, authToken *models.AuthToken) error
	AuthTokenUpdate(ctx context.Context, authToken *models.AuthToken) error
	AuthTokenRemove(ctx context.Context, id string) error
}

func (s *storeImpl) AuthTokenFind(ctx context.Context, id string) (*models.AuthToken, error) {
	f := new(models.AuthToken)
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

func (s *storeImpl) AuthTokenUpdate(ctx context.Context, authToken *models.AuthToken) error {
	_, err := s.db.
		NewUpdate().
		Model(authToken).
		WherePK().
		Returning("*").
		Exec(ctx)
	return err
}

func (s *storeImpl) AuthTokenCreate(ctx context.Context, authToken *models.AuthToken) error {
	_, err := s.db.
		NewInsert().
		Model(authToken).
		Exec(ctx)
	return err
}

func (s *storeImpl) AuthTokenRemove(ctx context.Context, id string) error {
	_, err := s.db.
		NewDelete().
		Model((*models.AuthToken)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
