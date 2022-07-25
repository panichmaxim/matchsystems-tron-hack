package models

import (
	"github.com/uptrace/bun"
)

type Category struct {
	bun.BaseModel `bun:"category,alias:category"`
	ID            int64  `bun:"id,pk,autoincrement" json:"id"`
	Name          string `bun:"name,notnull" json:"name"`
	Risk          int    `bun:"risk,notnull" json:"risk"`
	DescriptionRu string `bun:"description_ru,notnull" json:"description_ru"`
	DescriptionEn string `bun:"description_en,notnull" json:"description_en"`
}
