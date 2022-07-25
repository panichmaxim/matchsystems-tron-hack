package dataloader

import (
	"context"
	"github.com/rs/zerolog/log"
	"gitlab.com/rubin-dev/api/pkg/models"
	"gitlab.com/rubin-dev/api/pkg/service"
	"net/http"
	"time"
)

// go run github.com/vektah/dataloaden UserLoader int64 '*gitlab.com/rubin-dev/api/pkg/models.User'

const loadersKey = "dataloaders"

type Loaders struct {
	UserById UserLoader
}

const waitTimeout = 3 * time.Millisecond

func Middleware(svc service.Service, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), loadersKey, &Loaders{
			UserById: UserLoader{
				maxBatch: 100,
				wait:     waitTimeout,
				fetch: func(ids []int64) ([]*models.User, []error) {
					users, err := svc.UserListByID(r.Context(), ids)
					if err != nil {
						log.Err(err).Msg("dataloader UserListByID error")
						return nil, nil
					}

					return users, nil
				},
			},
		})

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
