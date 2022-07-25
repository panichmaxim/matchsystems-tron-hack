package models

import (
	"context"
	"github.com/uptrace/bun"
	"time"
)

type BillingRequest struct {
	bun.BaseModel `bun:"billing_request,alias:br"`
	ID            int64     `bun:"id,pk,autoincrement" json:"id"`
	UserID        int64     `bun:"user_id,notnull" json:"user_id"`
	Query         string    `bun:"query,notnull" json:"query"`
	Risk          int       `bun:"risk,notnull" json:"risk"`
	Category      string    `bun:"category,notnull" json:"category"`
	Network       string    `bun:"network,notnull" json:"network"`
	CreatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"created_at"`
}

func (u *BillingRequest) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		u.CreatedAt = time.Now()
	}
	return nil
}
