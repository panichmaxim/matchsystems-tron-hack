package database

import (
	"context"
	"io/fs"

	"github.com/rs/zerolog/log"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

func LoadMigrations(ctx context.Context, db *bun.DB, fsys fs.FS) error {
	migrations := migrate.NewMigrations(
		migrate.WithMigrationsDirectory("/migrations"),
	)
	if err := migrations.Discover(fsys); err != nil {
		return err
	}

	migrator := migrate.NewMigrator(
		db,
		migrations,
		migrate.WithTableName("migrations"),
		migrate.WithLocksTableName("migrations_lock"),
	)
	if err := migrator.Init(ctx); err != nil {
		return err
	}

	group, err := migrator.Migrate(ctx)
	if err != nil {
		return err
	}

	if group.ID == 0 {
		log.Info().Msg("there are no new migrations to run")
		return nil
	}

	log.Info().Msgf("migrated to %s", group)

	return nil
}
