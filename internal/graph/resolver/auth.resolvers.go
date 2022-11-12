package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"strings"

	"gitlab.com/rubin-dev/api/internal/graph/model"
	"gitlab.com/rubin-dev/api/pkg/models"
	"gitlab.com/rubin-dev/api/pkg/service"
	"gitlab.com/rubin-dev/api/pkg/validator"
)

func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.LoginResponse, error) {
	resp, err := r.svc.Login(ctx, &models.LoginRequest{
		Email:    strings.TrimSpace(input.Email),
		Password: input.Password,
	})
	if err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.LoginResponse{Errors: errs}, nil
		}

		return nil, err
	}

	return &model.LoginResponse{Jwt: resp}, err
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (*model.RefreshTokenResponse, error) {
	resp, err := r.svc.RefreshToken(ctx, &models.TokenRequest{Token: *input.Token})
	if err != nil {
		if errors.Is(err, service.ErrExpiredRefreshToken) {
			return nil, createGraphErr(ExpiredRefreshToken, err.Error())
		}
		if errors.Is(err, service.ErrCorruptedToken) {
			return nil, createGraphErr(CorruptedToken, err.Error())
		}

		if errs, ok := err.(validator.Errors); ok {
			return &model.RefreshTokenResponse{Errors: errs}, nil
		}

		return nil, err
	}

	return &model.RefreshTokenResponse{Result: resp}, nil
}

func (r *mutationResolver) Restore(ctx context.Context, input model.RestoreInput) (*model.RestoreResponse, error) {
	_, err := r.svc.Restore(ctx, &models.RestoreRequest{Email: *input.Email})
	if err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.RestoreResponse{Errors: errs}, nil
		}

		return nil, err
	}

	return &model.RestoreResponse{State: true}, nil
}

func (r *mutationResolver) RestoreCheck(ctx context.Context, input model.RestoreCheckInput) (*model.RestoreResponse, error) {
	err := r.svc.RestoreCheck(ctx, &models.TokenRequest{
		Token: *input.Token,
	})
	if err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.RestoreResponse{Errors: errs}, nil
		}

		return nil, err
	}

	return &model.RestoreResponse{State: true}, nil
}

func (r *mutationResolver) RestoreConfirm(ctx context.Context, input model.RestoreConfirmInput) (*model.RestoreConfirmResponse, error) {
	_, jwt, err := r.svc.RestoreConfirm(ctx, &models.RestoreConfirmRequest{
		Token:           *input.Token,
		Password:        *input.Password,
		PasswordConfirm: *input.PasswordConfirm,
	})
	if err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.RestoreConfirmResponse{Errors: errs}, nil
		}

		return nil, err
	}

	return &model.RestoreConfirmResponse{Jwt: jwt}, nil
}

func (r *mutationResolver) Registration(ctx context.Context, input model.RegistrationInput) (*model.RegistrationResponse, error) {
	_, err := r.svc.Registration(ctx, &models.RegistrationRequest{
		Email:           strings.TrimSpace(input.Email),
		Password:        input.Password,
		PasswordConfirm: input.PasswordConfirm,
		Name:            input.Name,
	})
	if err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.RegistrationResponse{Errors: errs}, nil
		}

		return nil, err
	}

	return &model.RegistrationResponse{State: true}, nil
}

func (r *mutationResolver) RegistrationConfirm(ctx context.Context, input model.RegistrationConfirmInput) (*model.RegistrationConfirmResponse, error) {
	_, jwt, err := r.svc.RegistrationConfirm(ctx, &models.TokenRequest{
		Token: *input.Token,
	})
	if err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.RegistrationConfirmResponse{Errors: errs}, nil
		}

		return nil, err
	}

	return &model.RegistrationConfirmResponse{Jwt: jwt}, nil
}

func (r *mutationResolver) ChangePassword(ctx context.Context, input model.ChangePasswordInput) (*model.ChangePasswordResponse, error) {
	u := WithUser(ctx)
	jwt, err := r.svc.ChangePassword(ctx, &models.ChangePasswordRequest{
		UserID:          u.ID,
		PasswordCurrent: input.PasswordCurrent,
		Password:        input.Password,
		PasswordConfirm: input.PasswordConfirm,
	})
	if err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.ChangePasswordResponse{Errors: errs}, nil
		}

		return nil, err
	}

	return &model.ChangePasswordResponse{Jwt: jwt}, nil
}

func (r *queryResolver) Me(ctx context.Context) (*models.User, error) {
	return WithUser(ctx), nil
}
