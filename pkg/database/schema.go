package database

import (
	"context"

	"github.com/uptrace/bun"
)

// CreateSchema вспомогательный метод для удаления / создания схемы
// базы данных. Используется в тестах или создании тестового стенда из
// фикстур.
func CreateSchema(ctx context.Context, db *bun.DB, models []interface{}) error {
	for _, m := range models {
		if _, err := db.NewDropTable().Model(m).IfExists().Exec(ctx); err != nil {
			return err
		}
		if _, err := db.NewCreateTable().Model(m).Exec(ctx); err != nil {
			return err
		}
	}

	return nil
}
