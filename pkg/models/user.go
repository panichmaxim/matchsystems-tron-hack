package models

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"users,alias:u"`
	ID            int64 `bun:"id,pk,autoincrement" json:"id"`
	// User fields
	Name        string   `bun:"name,notnull" json:"name"`
	Email       string   `bun:"email,notnull,unique" json:"-"`
	IsActive    bool     `bun:"is_active,notnull,default:false" json:"is_active"`
	Password    *string  `bun:"password,nullzero" json:"-"`
	Token       *string  `bun:"token,nullzero,unique" json:"-"`
	Permissions []string `bun:"permissions,nullzero,array" json:"permissions"`
	// Internal fields
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"updated_at"`
}

func (u *User) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		u.CreatedAt = time.Now()
		u.UpdatedAt = time.Now()
	case *bun.UpdateQuery:
		u.UpdatedAt = time.Now()
	}
	return nil
}
