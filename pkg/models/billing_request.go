package models

import (
	"context"
	"github.com/uptrace/bun"
	"time"
)

type BillingRequest struct {
	bun.BaseModel `bun:"billing_request,alias:billing_request"`
	ID            int64     `bun:"id,pk,autoincrement" json:"id"`
	UserID        int64     `bun:"user_id,notnull" json:"user_id"`
	Query         string    `bun:"query,notnull" json:"query"`
	Risk          float64   `bun:"risk" json:"risk"`
	IsReported    bool      `bun:"is_reported,notnull" json:"is_reported"`
	IsWallet      bool      `bun:"is_wallet,notnull" json:"is_wallet"`
	IsCalculated  bool      `bun:"is_calculated,notnull" json:"is_calculated"`
	Network       string    `bun:"network,notnull" json:"network"`
	Last          bool      `bun:"last,notnull" json:"last"`
	CreatedAt     time.Time `bun:"created_at,nullzero,notnull,type:timestamp" json:"created_at"`
}

func (u *BillingRequest) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		if u.CreatedAt.IsZero() {
			u.CreatedAt = time.Now()
		}
	}
	return nil
}
