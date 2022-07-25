package jwtoken

import (
	"github.com/cristalhq/jwt/v4"
	"time"
)

const AccessAudience = "user"
const RefreshAudience = "refresh"

type Service interface {
	Validate(token []byte) (*jwt.Token, error)
	CreateIssuedAt() time.Time
	CreateAccessExpiredAt() time.Time
	CreateRefreshExpiredAt() time.Time
	CreateAccessToken(jwtID string, issuedAt time.Time, expiredAt time.Time) (*jwt.Token, error)
	CreateRefreshToken(jwtID string, issuedAt time.Time, expiredAt time.Time) (*jwt.Token, error)
	Build(claim jwt.RegisteredClaims) (*jwt.Token, error)
}

type serviceImpl struct {
	secretKey       []byte
	accessDuration  time.Duration
	refreshDuration time.Duration
	signer          *jwt.HSAlg
	builder         *jwt.Builder
}

func NewService(
	secretKey []byte,
	accessDuration, refreshDuration time.Duration,
) (Service, error) {
	signer, err := jwt.NewSignerHS(jwt.HS256, secretKey)
	if err != nil {
		return nil, err
	}

	builder := jwt.NewBuilder(signer)

	return &serviceImpl{
		secretKey:       secretKey,
		accessDuration:  accessDuration,
		refreshDuration: refreshDuration,
		signer:          signer,
		builder:         builder,
	}, nil
}

func (s *serviceImpl) Validate(token []byte) (*jwt.Token, error) {
	newToken, err := jwt.Parse(token, s.signer)
	if err != nil {
		return nil, err
	}

	if err := s.signer.Verify(newToken); err != nil {
		return nil, err
	}

	return newToken, nil
}

func (s *serviceImpl) CreateIssuedAt() time.Time {
	return time.Now().UTC()
}

func (s *serviceImpl) CreateAccessExpiredAt() time.Time {
	return time.Now().Add(s.accessDuration).UTC()
}

func (s *serviceImpl) CreateRefreshExpiredAt() time.Time {
	return time.Now().Add(s.refreshDuration).UTC()
}

func (s *serviceImpl) CreateAccessToken(jwtID string, issuedAt time.Time, expiredAt time.Time) (*jwt.Token, error) {
	return s.builder.Build(s.createClaim(
		jwtID,
		[]string{AccessAudience},
		issuedAt,
		expiredAt,
	))
}

func (s *serviceImpl) createClaim(jwtID string, audience []string, issuedAt time.Time, expiredAt time.Time) jwt.RegisteredClaims {
	return jwt.RegisteredClaims{
		ID:        jwtID,
		Audience:  audience,
		IssuedAt:  &jwt.NumericDate{Time: issuedAt},
		ExpiresAt: &jwt.NumericDate{Time: expiredAt},
	}
}

func (s *serviceImpl) CreateRefreshToken(jwtID string, issuedAt time.Time, expiredAt time.Time) (*jwt.Token, error) {
	return s.builder.Build(s.createClaim(
		jwtID,
		[]string{RefreshAudience},
		issuedAt,
		expiredAt,
	))
}

func (s *serviceImpl) Build(claim jwt.RegisteredClaims) (*jwt.Token, error) {
	return s.builder.Build(claim)
}
