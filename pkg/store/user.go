package store

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"github.com/uptrace/bun"
	"gitlab.com/rubin-dev/api/pkg/models"
	"strings"
)

var _ UserStore = (*storeImpl)(nil)

type UserStore interface {
	UserFindByEmail(ctx context.Context, email string) (*models.User, error)
	UserFindByEmailActive(ctx context.Context, email string) (*models.User, error)
	UserFindByToken(ctx context.Context, token string) (*models.User, error)
	UserFindByID(ctx context.Context, id int64) (u *models.User, err error)
	UserRemoveByID(ctx context.Context, id int64) error
	UserCreate(ctx context.Context, u *models.User) error
	UserUpdate(ctx context.Context, u *models.User, columns ...string) error
	UserListAndPaginate(ctx context.Context, limit, offset int) (u []*models.User, count int, err error)
	UserListByID(ctx context.Context, ids []int64) ([]*models.User, error)
	UserFindSocialNetwork(ctx context.Context, network, identity string) (*models.User, error)
}

func (s *storeImpl) findByField(ctx context.Context, field string, value interface{}) (*models.User, error) {
	u := new(models.User)
	err := s.db.
		NewSelect().
		Model(u).
		Where(fmt.Sprintf("%s = ?", field), value).
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = nil
		}
		return nil, err
	}
	return u, nil
}

func (s *storeImpl) UserListByID(ctx context.Context, ids []int64) (users []*models.User, err error) {
	err = s.db.
		NewSelect().
		Model((*models.User)(nil)).
		Where("id IN (?)", bun.In(ids)).
		Scan(ctx, &users)
	return
}

func (s *storeImpl) UserFindByEmail(ctx context.Context, email string) (*models.User, error) {
	return s.findByField(ctx, "LOWER(email)", strings.ToLower(email))
}

func (s *storeImpl) UserFindByEmailActive(ctx context.Context, email string) (*models.User, error) {
	u := new(models.User)
	err := s.db.
		NewSelect().
		Model(u).
		Where("LOWER(email) = ? AND is_active = ?", strings.ToLower(email), true).
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = nil
		}
		return nil, err
	}
	return u, nil
}

func (s *storeImpl) UserCreate(ctx context.Context, u *models.User) (err error) {
	_, err = s.db.
		NewInsert().
		Model(u).
		Exec(ctx)
	return
}

func (s *storeImpl) UserUpdate(ctx context.Context, u *models.User, columns ...string) (err error) {
	_, err = s.db.
		NewUpdate().
		Model(u).
		WherePK().
		Returning("*").
		Column(columns...).
		Exec(ctx)
	return
}

func (s *storeImpl) UserRemoveByID(ctx context.Context, id int64) error {
	u := new(models.User)
	_, err := s.db.
		NewDelete().
		Model(u).
		Where("id = ?", id).
		Exec(ctx)
	return err
}

func (s *storeImpl) UserFindByID(ctx context.Context, id int64) (*models.User, error) {
	return s.findByField(ctx, "id", id)
}

func (s *storeImpl) UserFindByToken(ctx context.Context, token string) (*models.User, error) {
	return s.findByField(ctx, "token", token)
}

func (s *storeImpl) UserFindSocialNetwork(ctx context.Context, network, identity string) (*models.User, error) {
	u := new(models.User)
	err := s.db.
		NewSelect().
		Model(u).
		Where("social_network = ? AND social_identity = ?", network, identity).
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = nil
		}
		return nil, err
	}
	return u, nil
}

func (s *storeImpl) UserListAndPaginate(ctx context.Context, limit, offset int) (u []*models.User, count int, err error) {
	count, err = s.db.
		NewSelect().
		Model((*models.User)(nil)).
		Limit(limit).
		Offset(offset).
		ScanAndCount(ctx, &u)
	return
}
