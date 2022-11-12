package store

import (
	"context"
	"github.com/uptrace/bun"
	"gitlab.com/rubin-dev/api/pkg/database"
	"gitlab.com/rubin-dev/api/pkg/models"
	"io/fs"
)

var _ Store = (*storeImpl)(nil)

func NewSQLStore(db *bun.DB) Store {
	return &storeImpl{db}
}

type storeImpl struct {
	db *bun.DB
}

func (s *storeImpl) Health(ctx context.Context) error {
	return s.db.PingContext(ctx)
}

func (s *storeImpl) Close(ctx context.Context) error {
	return s.db.Close()
}

func (s *storeImpl) LoadFixtures(ctx context.Context, fixtureFS fs.FS, fixtures []string) error {
	if err := database.CreateSchema(ctx, s.db, models.GetModels()); err != nil {
		return err
	}
	if err := database.LoadFixtures(ctx, s.db, fixtureFS, fixtures); err != nil {
		return err
	}

	return nil
}

func (s *storeImpl) LoadMigrations(ctx context.Context, migrationFS fs.FS) error {
	return database.LoadMigrations(ctx, s.db, migrationFS)
}
