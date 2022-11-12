package models

import (
	"context"
	"github.com/uptrace/bun"
	"time"
)

type BillingRisk struct {
	bun.BaseModel    `bun:"billing_risk,alias:billing_risk"`
	ID               int64     `bun:"id,pk,autoincrement" json:"id"`
	BillingRequestID int64     `bun:"billing_request_id,notnull" json:"billing_request_id"`
	IsReported       bool      `bun:"is_reported,notnull" json:"is_reported"`
	IsWallet         bool      `bun:"is_wallet,notnull" json:"is_wallet"`
	IsCalculated     bool      `bun:"is_calculated,notnull" json:"is_calculated"`
	Risk             float64   `bun:"risk,notnull" json:"risk"`
	RiskRaw          float64   `bun:"risk_raw,notnull" json:"risk_raw"`
	Percent          float64   `bun:"percent,notnull" json:"percent"`
	PercentRaw       float64   `bun:"percent_raw,notnull" json:"percent_raw"`
	DirectoryID      int       `bun:"directory_id,notnull" json:"directory_id"`
	CategoryID       *int      `bun:"category_id,nullzero" json:"category_id"`
	Total            float64   `bun:"total,notnull" json:"total"`
	CreatedAt        time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"created_at"`
}

func (u *BillingRisk) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		if u.CreatedAt.IsZero() {
			u.CreatedAt = time.Now()
		}
	}
	return nil
}
