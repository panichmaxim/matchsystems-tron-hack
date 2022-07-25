package otlp

import (
	"context"
	"gitlab.com/rubin-dev/api/pkg/models"
	"gitlab.com/rubin-dev/api/pkg/service"
	"go.opentelemetry.io/otel/attribute"
)

var _ service.UserService = (*metricService)(nil)

func (m *metricService) UserHasPermissions(u *models.User, permissions []string) bool {
	return m.s.UserHasPermissions(u, permissions)
}

func (m *metricService) UserListByID(ctx context.Context, ids []int64) ([]*models.User, error) {
	ctx, t := m.tracer.Start(ctx, "service.UserListByID")
	t.SetAttributes(
		attribute.Int64Slice("identity", ids),
	)
	defer t.End()

	return m.s.UserListByID(ctx, ids)
}

func (m *metricService) UserHasPermission(user *models.User, permission string) bool {
	return m.s.UserHasPermission(user, permission)
}

func (m *metricService) UserFindByToken(ctx context.Context, token string) (*models.User, error) {
	ctx, t := m.tracer.Start(ctx, "service.UserFindByToken")
	defer t.End()

	return m.s.UserFindByToken(ctx, token)
}

func (m *metricService) UserUpdate(ctx context.Context, u *models.User, columns ...string) error {
	ctx, t := m.tracer.Start(ctx, "service.UserUpdate")
	t.SetAttributes(
		attribute.Int64("id", u.ID),
	)
	defer t.End()

	return m.s.UserUpdate(ctx, u, columns...)
}

func (m *metricService) UserProfileUpdate(ctx context.Context, u *models.User, req *models.UserProfileUpdateInput) error {
	ctx, t := m.tracer.Start(ctx, "service.UserProfileUpdate")
	t.SetAttributes(
		attribute.Int64("id", u.ID),
	)
	defer t.End()

	return m.s.UserProfileUpdate(ctx, u, req)
}

func (m *metricService) UserList(ctx context.Context, req *models.UserListRequest) ([]*models.User, int, error) {
	ctx, t := m.tracer.Start(ctx, "service.UserList")
	defer t.End()

	return m.s.UserList(ctx, req)
}

func (m *metricService) UserRemove(ctx context.Context, id int64) error {
	ctx, t := m.tracer.Start(ctx, "service.UserRemove")
	t.SetAttributes(
		attribute.Int64("id", id),
	)
	defer t.End()

	return m.s.UserRemove(ctx, id)
}

func (m *metricService) UserFindByEmail(ctx context.Context, email string) (*models.User, error) {
	ctx, t := m.tracer.Start(ctx, "service.UserFindByEmail")
	t.SetAttributes(
		attribute.String("email", email),
	)
	defer t.End()

	return m.s.UserFindByEmail(ctx, email)
}

func (m *metricService) UserFindByID(ctx context.Context, id int64) (*models.User, error) {
	ctx, t := m.tracer.Start(ctx, "service.UserFindByID")
	t.SetAttributes(
		attribute.Int64("id", id),
	)
	defer t.End()

	return m.s.UserFindByID(ctx, id)
}
