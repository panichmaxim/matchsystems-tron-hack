package service

import (
	"context"
	"github.com/rs/zerolog/log"
	"gitlab.com/rubin-dev/api/pkg/models"
)

var _ AccessRequestService = (*serviceImpl)(nil)

const accessRequestAdminUser = "panichmax@gmail.com"

type AccessRequestService interface {
	AccessRequest(ctx context.Context, user *models.User) error
}

func (s *serviceImpl) AccessRequest(ctx context.Context, user *models.User) error {
	if err := s.s.AccessRequest(ctx, user.ID); err != nil {
		return err
	}

	go func() {
		if err := s.mail.AccessRequest(ctx, accessRequestAdminUser, user.Name, user.Email); err != nil {
			log.Err(err).Msg("mail.AccessRequest")
		}
	}()

	return nil
}
