package models

import (
	"context"
	"github.com/uptrace/bun"
	"time"
)

// AccessRequest запрос доступа к платному функционалу
type AccessRequest struct {
	bun.BaseModel `bun:"access_requests,alias:ar"`
	ID            int64     `bun:"id,pk,autoincrement" json:"id"`
	UserID        int64     `bun:"user_id,notnull" json:"user_id"`
	CreatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"created_at"`
}

func (u *AccessRequest) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		u.CreatedAt = time.Now()
	}
	return nil
}
