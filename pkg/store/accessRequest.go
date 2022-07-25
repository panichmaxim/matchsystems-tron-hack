package store

import (
	"context"
	"gitlab.com/rubin-dev/api/pkg/models"
)

var _ AccessRequestStore = (*storeImpl)(nil)

type AccessRequestStore interface {
	AccessRequest(ctx context.Context, id int64) error
}

func (s *storeImpl) AccessRequest(ctx context.Context, id int64) error {
	ar := &models.AccessRequest{
		UserID: id,
	}
	_, err := s.db.NewInsert().Model(ar).Exec(ctx)
	return err
}
