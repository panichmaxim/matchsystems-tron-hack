package database

import (
	"context"
	"io/fs"
	"text/template"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dbfixture"
)

const hour24 = 24 * time.Hour
const day7 = 7 * hour24

// LoadFixtures вспомогательный метод для загрузки фикстур из переданной fs и списка
// yaml файлов. Поддерживается embed.FS.
func LoadFixtures(ctx context.Context, db *bun.DB, fsys fs.FS, files []string) error {
	funcMap := template.FuncMap{
		"now": func() string {
			return time.Now().Format(time.RFC3339Nano)
		},
		"prev_day": func() string {
			return time.Now().Add(-hour24).Format(time.RFC3339Nano)
		},
		"next_day": func() string {
			return time.Now().Add(hour24).Format(time.RFC3339Nano)
		},
		"next_week": func() string {
			return time.Now().Add(day7).Format(time.RFC3339Nano)
		},
		"prev_week": func() string {
			return time.Now().Add(-day7).Format(time.RFC3339Nano)
		},
	}
	return dbfixture.New(db, dbfixture.WithTemplateFuncs(funcMap)).Load(ctx, fsys, files...)
}
