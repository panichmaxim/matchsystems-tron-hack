package service

import (
	"context"
	"gitlab.com/rubin-dev/api/pkg/elastic"
	"gitlab.com/rubin-dev/api/pkg/jwtoken"
	"gitlab.com/rubin-dev/api/pkg/mailer"
	"gitlab.com/rubin-dev/api/pkg/neo4jstore/btc"
	"gitlab.com/rubin-dev/api/pkg/neo4jstore/eth"
	"gitlab.com/rubin-dev/api/pkg/neo4jstore/tron"
	"gitlab.com/rubin-dev/api/pkg/store"
	"gitlab.com/rubin-dev/api/pkg/validator"
)

var _ Service = (*serviceImpl)(nil)

func NewService(
	s store.Store,
	jwt jwtoken.Service,
	elk elastic.Client,
	mail mailer.Notify,
	btcneo btc.BtcStore,
	ethneo eth.EthStore,
	tronneo tron.TronStore,
	mockToken bool,
) Service {
	v := validator.NewValidationWithTracing(validator.NewValidation(s), "validation")
	return &serviceImpl{
		s:         s,
		v:         v,
		mail:      mail,
		jwt:       jwt,
		mockToken: mockToken,
		elk:       elk,
		btcneo:    btcneo,
		ethneo:    ethneo,
		tronneo:   tronneo,
	}
}

type serviceImpl struct {
	s         store.Store
	v         validator.Validation
	mockToken bool
	mail      mailer.Notify
	jwt       jwtoken.Service
	elk       elastic.Client
	btcneo    btc.BtcStore
	ethneo    eth.EthStore
	tronneo   tron.TronStore
}

func (s *serviceImpl) Health(ctx context.Context) error {
	if err := s.s.Health(ctx); err != nil {
		return err
	}
	if err := s.btcneo.Health(ctx); err != nil {
		return err
	}
	if err := s.ethneo.Health(ctx); err != nil {
		return err
	}
	if err := s.tronneo.Health(ctx); err != nil {
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
