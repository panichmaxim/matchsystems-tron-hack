package models

import (
	"context"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type BillingKey struct {
	bun.BaseModel `bun:"billing_key,alias:bak"`
	ID            int64     `bun:"id,pk,autoincrement" json:"id"`
	Key           string    `bun:"key,unique,notnull" json:"key"`
	UserID        int64     `bun:"user_id,notnull" json:"user_id"`
	BillingID     int64     `bun:"billing_id,notnull" json:"billing_id"`
	CreatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"created_at"`
}

func (u *BillingKey) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		u.CreatedAt = time.Now()
		if len(u.Key) == 0 {
			u.Key = uuid.NewString()
		}
	}
	return nil
}
