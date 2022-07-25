package service

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/cristalhq/jwt/v4"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"gitlab.com/rubin-dev/api/pkg/jwtoken"
	"gitlab.com/rubin-dev/api/pkg/models"
	"gitlab.com/rubin-dev/api/pkg/password"
	"time"
)

var _ AuthService = (*serviceImpl)(nil)

type AuthService interface {
	GetUserFromRequest(ctx context.Context, accessToken string) (*models.User, error)
	GetUserFromRefreshToken(ctx context.Context, refreshToken string) (*models.User, *models.Jwt, error)
	Login(ctx context.Context, req *models.LoginRequest) (*models.Jwt, error)
	ForceLogin(ctx context.Context, id int64) (*models.Jwt, error)
	RefreshToken(ctx context.Context, req *models.TokenRequest) (*models.Jwt, error)
	Registration(ctx context.Context, req *models.RegistrationRequest) (*models.User, error)
	RegistrationConfirm(ctx context.Context, req *models.TokenRequest) (*models.User, *models.Jwt, error)
	Restore(ctx context.Context, req *models.RestoreRequest) (*models.User, error)
	RestoreCheck(ctx context.Context, req *models.TokenRequest) error
	RestoreConfirm(ctx context.Context, req *models.RestoreConfirmRequest) (*models.User, *models.Jwt, error)
	ChangePassword(ctx context.Context, req *models.ChangePasswordRequest) (*models.Jwt, error)
}

type ContextKeyType string

const ContextUserKey ContextKeyType = "user"

var ErrInvalidAccessToken = errors.New("invalid access token")
var ErrInvalidRefreshToken = errors.New("invalid refresh token")
var ErrExpiredAccessToken = errors.New("expired access token")
var ErrExpiredRefreshToken = errors.New("expired refresh token")
var ErrInvalidAudience = errors.New("invalid audience")
var ErrCorruptedToken = errors.New("token corrupted")

func (s *serviceImpl) GetUserFromRequest(ctx context.Context, accessToken string) (*models.User, error) {
	token, err := s.jwt.Validate([]byte(accessToken))
	if err != nil {
		log.Debug().Err(err).Msg("invalid access token")
		return nil, ErrInvalidAccessToken
	}

	var claim jwt.RegisteredClaims
	if err := token.DecodeClaims(&claim); err != nil {
		return nil, err
	}

	if !claim.IsForAudience(jwtoken.AccessAudience) {
		return nil, ErrInvalidAudience
	}

	if !claim.IsValidAt(time.Now()) {
		return nil, ErrExpiredAccessToken
	}

	authToken, err := s.s.AuthTokenFind(ctx, claim.ID)
	if err != nil {
		return nil, err
	}
	if authToken == nil {
		return nil, ErrCorruptedToken
	}
	if time.Now().After(authToken.AccessExpiredAt) || time.Now().After(claim.ExpiresAt.Time) {
		return nil, ErrExpiredAccessToken
	}

	return s.UserFindByID(ctx, authToken.UserID)
}

func (s *serviceImpl) GetUserFromRefreshToken(ctx context.Context, refreshToken string) (*models.User, *models.Jwt, error) {
	resp, err := s.RefreshToken(ctx, &models.TokenRequest{Token: refreshToken})
	if err != nil {
		return nil, nil, errors.Wrap(err, "refresh token")
	}

	u, err := s.GetUserFromRequest(ctx, resp.AccessToken)
	if err != nil {
		return nil, nil, err
	}

	return u, resp, nil
}

func (s *serviceImpl) GetUserFromContext(ctx context.Context) *models.User {
	u, ok := ctx.Value(ContextUserKey).(*models.User)
	if !ok {
		return nil
	}

	return u
}

func (s *serviceImpl) createToken(email string) string {
	if s.mockToken {
		return fmt.Sprintf("%x", md5.Sum([]byte(email)))
	}

	return uuid.New().String()
}

func (s *serviceImpl) createJwt(ctx context.Context, user *models.User) (*models.AuthToken, error) {
	id := uuid.New()

	issuedAt := s.jwt.CreateIssuedAt()

	accessExpiredAt := s.jwt.CreateAccessExpiredAt()
	accessToken, err := s.jwt.CreateAccessToken(id.String(), issuedAt, accessExpiredAt)
	if err != nil {
		return nil, err
	}

	refreshExpiredAt := s.jwt.CreateRefreshExpiredAt()
	refreshToken, err := s.jwt.CreateRefreshToken(id.String(), issuedAt, refreshExpiredAt)
	if err != nil {
		return nil, err
	}

	authToken := &models.AuthToken{
		ID:               id,
		UserID:           user.ID,
		AccessToken:      accessToken.String(),
		AccessExpiredAt:  accessExpiredAt,
		RefreshToken:     refreshToken.String(),
		RefreshExpiredAt: refreshExpiredAt,
		IssuedAt:         issuedAt,
	}
	if err := s.s.AuthTokenCreate(ctx, authToken); err != nil {
		return nil, err
	}

	return authToken, nil
}

func (s *serviceImpl) refreshJwtToken(ctx context.Context, authToken *models.AuthToken, user *models.User) error {
	issuedAt := s.jwt.CreateIssuedAt()

	accessExpiredAt := s.jwt.CreateAccessExpiredAt()
	accessToken, err := s.jwt.CreateAccessToken(authToken.ID.String(), issuedAt, accessExpiredAt)
	if err != nil {
		return err
	}

	refreshExpiredAt := s.jwt.CreateRefreshExpiredAt()
	refreshToken, err := s.jwt.CreateRefreshToken(authToken.ID.String(), issuedAt, refreshExpiredAt)
	if err != nil {
		return err
	}

	authToken.AccessToken = accessToken.String()
	authToken.AccessExpiredAt = accessExpiredAt
	authToken.RefreshToken = refreshToken.String()
	authToken.RefreshExpiredAt = refreshExpiredAt
	authToken.IssuedAt = issuedAt

	return s.s.AuthTokenUpdate(ctx, authToken)
}

func (s *serviceImpl) Login(ctx context.Context, req *models.LoginRequest) (*models.Jwt, error) {
	if err := s.v.AuthLogin(ctx, req); err != nil {
		return nil, err
	}

	user, err := s.s.UserFindByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	authToken, err := s.createJwt(ctx, user)
	if err != nil {
		return nil, err
	}

	return &models.Jwt{
		ID:           user.ID,
		Permissions:  user.Permissions,
		AccessToken:  authToken.AccessToken,
		RefreshToken: authToken.RefreshToken,
	}, nil
}

func (s *serviceImpl) ForceLogin(ctx context.Context, id int64) (*models.Jwt, error) {
	user, err := s.s.UserFindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	authToken, err := s.createJwt(ctx, user)
	if err != nil {
		return nil, err
	}

	return &models.Jwt{
		ID:           user.ID,
		Permissions:  user.Permissions,
		AccessToken:  authToken.AccessToken,
		RefreshToken: authToken.RefreshToken,
	}, nil
}

func (s *serviceImpl) Registration(ctx context.Context, req *models.RegistrationRequest) (*models.User, error) {
	if err := s.v.AuthRegistration(ctx, req); err != nil {
		return nil, err
	}

	hashPassword, err := password.HashPassword([]byte(req.Password))
	if err != nil {
		return nil, err
	}

	t := s.createToken(req.Email)
	p := string(hashPassword)
	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: &p,
		IsActive: false,
		Token:    &t,
	}

	err = s.s.UserCreate(ctx, user)
	if err != nil {
		return nil, err
	}

	if err := s.mail.Registration(ctx, user.Email, t); err != nil {
		log.Err(err).Msg("mail.Registration")
	}

	return user, nil
}

func (s *serviceImpl) RegistrationConfirm(ctx context.Context, req *models.TokenRequest) (*models.User, *models.Jwt, error) {
	if err := s.v.AuthRegistrationConfirm(ctx, req); err != nil {
		return nil, nil, err
	}

	user, err := s.s.UserFindByToken(ctx, req.Token)
	if err != nil {
		return nil, nil, err
	}

	user.Token = nil
	user.IsActive = true

	err = s.s.UserUpdate(ctx, user, "token", "is_active")
	if err != nil {
		return nil, nil, err
	}

	authToken, err := s.createJwt(ctx, user)
	if err != nil {
		return nil, nil, err
	}

	if err := s.mail.RegistrationConfirm(ctx, user.Email); err != nil {
		log.Err(err).Msg("mail.RegistrationConfirm")
	}

	return user, &models.Jwt{
		ID:           user.ID,
		Permissions:  user.Permissions,
		RefreshToken: authToken.RefreshToken,
		AccessToken:  authToken.AccessToken,
	}, nil
}

func (s *serviceImpl) RefreshToken(ctx context.Context, req *models.TokenRequest) (*models.Jwt, error) {
	token, err := s.jwt.Validate([]byte(req.Token))
	if err != nil {
		return nil, ErrInvalidRefreshToken
	}

	var claim jwt.RegisteredClaims
	if err := json.Unmarshal(token.Claims(), &claim); err != nil {
		return nil, err
	}

	if !claim.IsForAudience(jwtoken.RefreshAudience) {
		return nil, ErrInvalidAudience
	}

	if !claim.IsValidAt(time.Now()) {
		return nil, ErrExpiredRefreshToken
	}

	authToken, err := s.s.AuthTokenFind(ctx, claim.ID)
	if err != nil {
		return nil, err
	}
	if authToken == nil {
		return nil, ErrCorruptedToken
	}
	if time.Now().After(authToken.RefreshExpiredAt) || time.Now().After(claim.ExpiresAt.Time) {
		return nil, ErrExpiredRefreshToken
	}

	user, err := s.s.UserFindByID(ctx, authToken.UserID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	if err := s.refreshJwtToken(ctx, authToken, user); err != nil {
		return nil, err
	}

	return &models.Jwt{
		ID:           user.ID,
		Permissions:  user.Permissions,
		AccessToken:  authToken.AccessToken,
		RefreshToken: authToken.RefreshToken,
	}, nil
}

func (s *serviceImpl) Restore(ctx context.Context, req *models.RestoreRequest) (*models.User, error) {
	if err := s.v.AuthRestore(ctx, req); err != nil {
		return nil, err
	}

	user, err := s.s.UserFindByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	if user.Token == nil {
		t := s.createToken(req.Email)
		user.Token = &t
		if err = s.s.UserUpdate(ctx, user, "token"); err != nil {
			return nil, err
		}
	}

	if err := s.mail.Restore(ctx, user.Email, *user.Token); err != nil {
		log.Err(err).Msg("mail.Restore")
	}

	return user, nil
}

func (s *serviceImpl) RestoreCheck(ctx context.Context, req *models.TokenRequest) error {
	return s.v.AuthRestoreCheck(ctx, req)
}

func (s *serviceImpl) RestoreConfirm(ctx context.Context, req *models.RestoreConfirmRequest) (*models.User, *models.Jwt, error) {
	if err := s.v.AuthRestoreConfirm(ctx, req); err != nil {
		return nil, nil, err
	}

	user, err := s.s.UserFindByToken(ctx, req.Token)
	if err != nil {
		return nil, nil, err
	}

	user.Token = nil
	user.IsActive = true

	passwordHash, err := password.HashPassword([]byte(req.Password))
	if err != nil {
		return nil, nil, err
	}

	p := string(passwordHash)
	user.Password = &p

	if err := s.s.UserUpdate(ctx, user, "token", "is_active", "password"); err != nil {
		return nil, nil, err
	}

	authToken, err := s.createJwt(ctx, user)
	if err != nil {
		return nil, nil, err
	}

	if err := s.mail.RestoreConfirm(ctx, user.Email); err != nil {
		log.Err(err).Msg("mail.RestoreConfirm")
	}

	return user, &models.Jwt{
		ID:           user.ID,
		Permissions:  user.Permissions,
		RefreshToken: authToken.RefreshToken,
		AccessToken:  authToken.AccessToken,
	}, nil
}

func (s *serviceImpl) ChangePassword(ctx context.Context, req *models.ChangePasswordRequest) (*models.Jwt, error) {
	if err := s.v.AuthChangePassword(ctx, req); err != nil {
		return nil, err
	}

	user, err := s.s.UserFindByID(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	passwordHash, err := password.HashPassword([]byte(req.Password))
	if err != nil {
		return nil, err
	}

	p := string(passwordHash)
	user.Password = &p

	if err := s.s.UserUpdate(ctx, user, "password"); err != nil {
		return nil, err
	}

	authToken, err := s.createJwt(ctx, user)
	if err != nil {
		return nil, err
	}

	return &models.Jwt{
		ID:           user.ID,
		Permissions:  user.Permissions,
		RefreshToken: authToken.RefreshToken,
		AccessToken:  authToken.AccessToken,
	}, nil
}
