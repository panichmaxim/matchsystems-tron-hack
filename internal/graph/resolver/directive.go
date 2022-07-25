package resolver

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"gitlab.com/rubin-dev/api/pkg/jwtoken"
	"gitlab.com/rubin-dev/api/pkg/models"
	"gitlab.com/rubin-dev/api/pkg/service"
	"go.opentelemetry.io/otel/trace"
)

const (
	AccessDenied        = "PERMISSION_DENIED"
	ExpiredToken        = "EXPIRED_TOKEN"
	ExpiredRefreshToken = "EXPIRED_REFRESH_TOKEN"
	CorruptedToken      = "CORRUPTED_TOKEN"
)

type contextKey struct {
	fieldName string
}

var UserCtxKey = &contextKey{"user"}

func WithUser(ctx context.Context) *models.User {
	return ctx.Value(UserCtxKey).(*models.User)
}

func NewAuthDirective(
	svc service.Service,
	tracer trace.Tracer,
) *AuthDirective {
	return &AuthDirective{
		svc:    svc,
		tracer: tracer,
	}
}

type AuthDirective struct {
	svc    service.Service
	tracer trace.Tracer
}

func createGraphErr(msg, extra string) *gqlerror.Error {
	return &gqlerror.Error{
		Message:    msg,
		Extensions: map[string]interface{}{"message": extra},
	}
}

func (a *AuthDirective) Auth(ctx context.Context, obj interface{}, next graphql.Resolver, permissions []string) (res interface{}, err error) {
	ctx, t := a.tracer.Start(ctx, "AuthDirective.Auth")
	defer t.End()

	authorizationKey := jwtoken.WithAuthorizationContext(ctx)
	if len(authorizationKey) == 0 {
		return nil, createGraphErr(AccessDenied, "empty or missing authorization header")
	}

	user, err := a.svc.GetUserFromRequest(ctx, authorizationKey)
	if err != nil {
		if errors.Is(err, service.ErrExpiredAccessToken) {
			return nil, createGraphErr(ExpiredToken, err.Error())
		}
		if errors.Is(err, service.ErrCorruptedToken) {
			return nil, createGraphErr(CorruptedToken, err.Error())
		}
		return nil, createGraphErr(AccessDenied, err.Error())
	}

	if !a.svc.UserHasPermissions(user, permissions) {
		return nil, createGraphErr(AccessDenied, "permission denied")
	}

	return next(context.WithValue(ctx, UserCtxKey, user))
}

func (a *AuthDirective) AuthRefresh(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	ctx, t := a.tracer.Start(ctx, "AuthDirective.AuthRefresh")
	defer t.End()

	authorizationKey := jwtoken.WithAuthorizationContext(ctx)
	if len(authorizationKey) == 0 {
		return nil, createGraphErr(AccessDenied, "empty or missing authorization header")
	}

	user, _, err := a.svc.GetUserFromRefreshToken(ctx, authorizationKey)
	if err != nil {
		if errors.Is(err, service.ErrExpiredAccessToken) {
			return nil, createGraphErr(ExpiredToken, err.Error())
		}
		if errors.Is(err, service.ErrExpiredRefreshToken) {
			return nil, createGraphErr(ExpiredRefreshToken, err.Error())
		}
		if errors.Is(err, service.ErrCorruptedToken) {
			return nil, createGraphErr(CorruptedToken, err.Error())
		}

		return nil, createGraphErr(AccessDenied, err.Error())
	}

	return next(context.WithValue(ctx, UserCtxKey, user))
}
