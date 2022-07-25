package models

import (
	"context"
	"github.com/uptrace/bun"
	"time"
)

type BillingPacket struct {
	bun.BaseModel `bun:"billing_packet,alias:br"`
	ID            int64     `bun:"id,pk,autoincrement" json:"id"`
	UserID        int64     `bun:"user_id,notnull" json:"user_id"`
	Requests      int       `bun:"requests,notnull" json:"requests"`
	CreatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"created_at"`
}

func (u *BillingPacket) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		u.CreatedAt = time.Now()
	}
	return nil
}
