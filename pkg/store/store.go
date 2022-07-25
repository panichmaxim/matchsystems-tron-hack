package store

import (
	"context"
	"io/fs"
)

type Store interface {
	UserStore
	AuthTokenStore
	AccessRequestStore
	BillingStore
	CategoryStore

	Close(ctx context.Context) error
	LoadMigrations(ctx context.Context, migrationFS fs.FS) error
	LoadFixtures(ctx context.Context, fixtureFS fs.FS, fixtures []string) error
	Health(ctx context.Context) error
}
