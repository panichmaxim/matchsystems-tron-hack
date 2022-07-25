package otlp

import (
	"context"
	"gitlab.com/rubin-dev/api/pkg/models"
	"gitlab.com/rubin-dev/api/pkg/service"
	"go.opentelemetry.io/otel/attribute"
)

var _ service.AuthService = (*metricService)(nil)

func (m *metricService) GetUserFromRefreshToken(ctx context.Context, refreshToken string) (*models.User, *models.Jwt, error) {
	ctx, t := m.tracer.Start(ctx, "service.GetAndRefreshUserFromRequest")
	defer t.End()

	return m.s.GetUserFromRefreshToken(ctx, refreshToken)
}

func (m *metricService) GetUserFromRequest(ctx context.Context, accessToken string) (*models.User, error) {
	ctx, t := m.tracer.Start(ctx, "service.GetUserFromRequest")
	defer t.End()

	return m.s.GetUserFromRequest(ctx, accessToken)
}

func (m *metricService) Login(ctx context.Context, req *models.LoginRequest) (*models.Jwt, error) {
	ctx, t := m.tracer.Start(ctx, "service.Login")
	t.SetAttributes(
		attribute.String("email", req.Email),
	)
	defer t.End()

	return m.s.Login(ctx, req)
}

func (m *metricService) ForceLogin(ctx context.Context, id int64) (*models.Jwt, error) {
	ctx, t := m.tracer.Start(ctx, "service.ForceLogin")
	t.SetAttributes(
		attribute.Int64("id", id),
	)
	defer t.End()

	return m.s.ForceLogin(ctx, id)
}

func (m *metricService) RefreshToken(ctx context.Context, req *models.TokenRequest) (*models.Jwt, error) {
	ctx, t := m.tracer.Start(ctx, "service.RefreshToken")
	defer t.End()

	return m.s.RefreshToken(ctx, req)
}

func (m *metricService) Registration(ctx context.Context, req *models.RegistrationRequest) (*models.User, error) {
	ctx, t := m.tracer.Start(ctx, "service.Registration")
	t.SetAttributes(
		attribute.String("email", req.Email),
	)
	defer t.End()

	return m.s.Registration(ctx, req)
}

func (m *metricService) RegistrationConfirm(ctx context.Context, req *models.TokenRequest) (*models.User, *models.Jwt, error) {
	ctx, t := m.tracer.Start(ctx, "service.RegistrationConfirm")
	defer t.End()

	return m.s.RegistrationConfirm(ctx, req)
}

func (m *metricService) Restore(ctx context.Context, req *models.RestoreRequest) (*models.User, error) {
	ctx, t := m.tracer.Start(ctx, "service.Restore")
	t.SetAttributes(
		attribute.String("email", req.Email),
	)
	defer t.End()

	return m.s.Restore(ctx, req)
}

func (m *metricService) RestoreCheck(ctx context.Context, req *models.TokenRequest) error {
	ctx, t := m.tracer.Start(ctx, "service.RestoreCheck")
	defer t.End()

	return m.s.RestoreCheck(ctx, req)
}

func (m *metricService) RestoreConfirm(ctx context.Context, req *models.RestoreConfirmRequest) (*models.User, *models.Jwt, error) {
	ctx, t := m.tracer.Start(ctx, "service.RestoreConfirm")
	defer t.End()

	return m.s.RestoreConfirm(ctx, req)
}

func (m *metricService) ChangePassword(ctx context.Context, req *models.ChangePasswordRequest) (*models.Jwt, error) {
	ctx, t := m.tracer.Start(ctx, "service.ChangePassword")
	t.SetAttributes(
		attribute.Int64("id", req.UserID),
	)
	defer t.End()

	return m.s.ChangePassword(ctx, req)
}
