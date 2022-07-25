package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type AuthToken struct {
	bun.BaseModel    `bun:"auth_token,alias:at"`
	ID               uuid.UUID `bun:"type:uuid,pk"`
	UserID           int64     `bun:"user_id,notnull" json:"user_id"`
	AccessToken      string    `bun:"access_token,unique" json:"access_token"`
	AccessExpiredAt  time.Time `bun:",nullzero,notnull" json:"access_expired_at"`
	RefreshToken     string    `bun:"refresh_token,unique" json:"refresh_token"`
	RefreshExpiredAt time.Time `bun:",nullzero,notnull" json:"refresh_expired_at"`
	IssuedAt         time.Time `bun:",nullzero,notnull" json:"issued_at"`
}
