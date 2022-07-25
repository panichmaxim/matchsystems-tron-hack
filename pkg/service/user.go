package service

import (
	"context"
	"gitlab.com/rubin-dev/api/pkg/models"
)

var _ UserService = (*serviceImpl)(nil)

type UserService interface {
	UserListByID(ctx context.Context, ids []int64) ([]*models.User, error)
	UserFindByToken(ctx context.Context, token string) (*models.User, error)
	UserUpdate(ctx context.Context, u *models.User, columns ...string) error
	UserProfileUpdate(ctx context.Context, u *models.User, req *models.UserProfileUpdateInput) error
	UserList(ctx context.Context, req *models.UserListRequest) ([]*models.User, int, error)
	UserRemove(ctx context.Context, id int64) error
	UserHasPermission(u *models.User, permission string) bool
	UserHasPermissions(u *models.User, permissions []string) bool
	UserFindByEmail(ctx context.Context, email string) (*models.User, error)
	UserFindByID(ctx context.Context, id int64) (*models.User, error)
}

func (s *serviceImpl) UserHasPermissions(user *models.User, permissions []string) bool {
	if len(permissions) == 0 {
		return true
	}

	for _, r := range permissions {
		if !s.UserHasPermission(user, r) {
			return false
		}
	}

	return true
}

func (s *serviceImpl) UserHasPermission(user *models.User, permission string) bool {
	if user.Permissions == nil {
		return false
	}

	for _, v := range user.Permissions {
		if permission == v {
			return true
		}
	}

	return false
}

func (s *serviceImpl) UserList(ctx context.Context, req *models.UserListRequest) ([]*models.User, int, error) {
	limit, offset := buildLimitOffset(req.Page, req.PageSize)
	return s.s.UserListAndPaginate(ctx, limit, offset)
}

func (s *serviceImpl) UserFindByID(ctx context.Context, id int64) (*models.User, error) {
	return s.s.UserFindByID(ctx, id)
}

func (s *serviceImpl) UserRemove(ctx context.Context, id int64) error {
	return s.s.UserRemoveByID(ctx, id)
}

func (s *serviceImpl) UserFindByEmail(ctx context.Context, email string) (*models.User, error) {
	return s.s.UserFindByEmail(ctx, email)
}

func (s *serviceImpl) UserFindByToken(ctx context.Context, token string) (*models.User, error) {
	return s.s.UserFindByToken(ctx, token)
}

func (s *serviceImpl) UserUpdate(ctx context.Context, u *models.User, columns ...string) error {
	return s.s.UserUpdate(ctx, u, columns...)
}

func (s *serviceImpl) UserListByID(ctx context.Context, ids []int64) ([]*models.User, error) {
	return s.s.UserListByID(ctx, ids)
}

func (s *serviceImpl) UserProfileUpdate(ctx context.Context, u *models.User, req *models.UserProfileUpdateInput) error {
	if err := s.v.UserProfileUpdate(ctx, req); err != nil {
		return err
	}
	u.Name = *req.Name
	return s.s.UserUpdate(ctx, u, "name")
}
