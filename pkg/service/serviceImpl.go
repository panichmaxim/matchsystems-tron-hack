package service

import (
	"context"
	"gitlab.com/rubin-dev/api/pkg/btcstore"
	"gitlab.com/rubin-dev/api/pkg/elastic"
	"gitlab.com/rubin-dev/api/pkg/ethstore"
	"gitlab.com/rubin-dev/api/pkg/jwtoken"
	"gitlab.com/rubin-dev/api/pkg/mailer"
	"gitlab.com/rubin-dev/api/pkg/store"
	"gitlab.com/rubin-dev/api/pkg/validator"
)

var _ Service = (*serviceImpl)(nil)

func NewService(
	s store.Store,
	jwt jwtoken.Service,
	elk elastic.Client,
	mail mailer.Notify,
	btcneo btcstore.Store,
	ethneo ethstore.Store,
	mockToken bool,
) Service {
	return &serviceImpl{
		s:         s,
		v:         validator.NewValidation(s),
		mail:      mail,
		jwt:       jwt,
		mockToken: mockToken,
		elk:       elk,
		btcneo:    btcneo,
		ethneo:    ethneo,
	}
}

type serviceImpl struct {
	s         store.Store
	v         validator.Validation
	mockToken bool
	mail      mailer.Notify
	jwt       jwtoken.Service
	elk       elastic.Client
	btcneo    btcstore.Store
	ethneo    ethstore.Store
}

func (s *serviceImpl) Health(ctx context.Context) error {
	if err := s.s.Health(ctx); err != nil {
		return err
	}
	return nil
}

func (s *serviceImpl) Close(ctx context.Context) error {
	if err := s.s.Close(ctx); err != nil {
		return err
	}

	return nil
}
