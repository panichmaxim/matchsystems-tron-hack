package service

import (
	"context"
	"github.com/rs/zerolog/log"
	"gitlab.com/rubin-dev/api/pkg/models"
	"sync"
)

var _ AccessRequestService = (*serviceImpl)(nil)

const accessRequestAdmin1User = "panichmax@gmail.com"
const accessRequestAdmin2User = "maxfalaleev1@gmail.com"

type AccessRequestService interface {
	AccessRequest(ctx context.Context, user *models.User) error
}

func (s *serviceImpl) AccessRequest(ctx context.Context, user *models.User) error {
	if err := s.s.AccessRequest(ctx, user.ID); err != nil {
		return err
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		if err := s.mail.AccessRequest(ctx, accessRequestAdmin1User, user.Name, user.Email); err != nil {
			log.Err(err).Msg("mail.AccessRequest")
		}
	}()

	go func() {
		defer wg.Done()
		if err := s.mail.AccessRequest(ctx, accessRequestAdmin2User, user.Name, user.Email); err != nil {
			log.Err(err).Msg("mail.AccessRequest")
		}
	}()

	wg.Wait()

	return nil
}
