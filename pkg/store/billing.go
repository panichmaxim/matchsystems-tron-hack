package store

import (
	"context"
	"database/sql"
	"errors"
	"gitlab.com/rubin-dev/api/pkg/models"
)

var ErrInsufficientBalance = errors.New("insufficient balance")

var _ BillingStore = (*storeImpl)(nil)

type BillingStore interface {
	BillingGetOrCreate(ctx context.Context, userID int64) (*models.Billing, error)
	BillingAddRequests(ctx context.Context, userID int64, requests int) (*models.BillingPacket, error)
	BillingRegisterRequest(ctx context.Context, userID int64, query string, risk int, category string, network string) (*models.Billing, error)
	BillingSmartRegisterRequest(ctx context.Context, userID int64, query string, risk int, category string, network string) (*models.Billing, error)
	BillingFindApiKey(ctx context.Context, apiKey string) (*models.BillingKey, error)
	BillingCreateApiKey(ctx context.Context, userID int64) (*models.BillingKey, error)
	BillingHistoryList(ctx context.Context, userID int64, limit, offset int) ([]*models.BillingRequest, int, error)
	BillingKeyList(ctx context.Context, userID int64, limit, offset int) ([]*models.BillingKey, int, error)
	BillingKeyRemove(ctx context.Context, userID, id int64) error
	BillingHistoryRequestsCount(ctx context.Context, userID int64) (int, error)
}

func (s *storeImpl) BillingHistoryRequestsCount(ctx context.Context, userID int64) (int, error) {
	var count int
	err := s.db.
		NewSelect().
		Model((*models.BillingRequest)(nil)).
		ColumnExpr("count(id)").
		Where("user_id = ?", userID).
		Scan(ctx, &count)
	return count, err
}

func (s *storeImpl) BillingKeyRemove(ctx context.Context, userID, id int64) error {
	_, err := s.db.
		NewDelete().
		Model((*models.BillingKey)(nil)).
		Where("user_id = ? AND id = ?", userID, id).
		Exec(ctx)
	return err
}

func (s *storeImpl) BillingKeyList(ctx context.Context, userID int64, limit, offset int) ([]*models.BillingKey, int, error) {
	var u []*models.BillingKey
	count, err := s.db.
		NewSelect().
		Model((*models.BillingKey)(nil)).
		Where("user_id = ?", userID).
		Order("id DESC").
		Limit(limit).
		Offset(offset).
		ScanAndCount(ctx, &u)
	return u, count, err
}

func (s *storeImpl) BillingHistoryList(ctx context.Context, userID int64, limit, offset int) ([]*models.BillingRequest, int, error) {
	var u []*models.BillingRequest
	count, err := s.db.
		NewSelect().
		Model((*models.BillingRequest)(nil)).
		Where("user_id = ?", userID).
		Order("id DESC").
		Limit(limit).
		Offset(offset).
		ScanAndCount(ctx, &u)
	return u, count, err
}

func (s *storeImpl) BillingFindApiKey(ctx context.Context, apiKey string) (*models.BillingKey, error) {
	billingApiKey := new(models.BillingKey)
	err := s.db.
		NewSelect().
		Model((*models.BillingKey)(nil)).
		Where("key = ?", apiKey).
		Scan(ctx, billingApiKey)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return billingApiKey, nil
}

func (s *storeImpl) BillingCreateApiKey(ctx context.Context, userID int64) (*models.BillingKey, error) {
	billing, err := s.BillingGetOrCreate(ctx, userID)
	if err != nil {
		return nil, err
	}
	billingApiKey := &models.BillingKey{
		UserID:    userID,
		BillingID: billing.ID,
	}
	_, err = s.db.
		NewInsert().
		Model(billingApiKey).
		Exec(ctx)
	return billingApiKey, nil
}

func (s *storeImpl) BillingGetOrCreate(ctx context.Context, userID int64) (*models.Billing, error) {
	billing := new(models.Billing)
	err := s.db.NewSelect().Model((*models.Billing)(nil)).Where("user_id = ?", userID).Scan(ctx, billing)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}

		billing = &models.Billing{UserID: userID}
		_, err := s.db.
			NewInsert().
			Model(billing).
			Exec(ctx)
		if err != nil {
			return nil, err
		}

		return billing, nil
	}

	return billing, nil
}

func (s *storeImpl) BillingAddRequests(ctx context.Context, userID int64, requests int) (*models.BillingPacket, error) {
	billing, err := s.BillingGetOrCreate(ctx, userID)
	if err != nil {
		return nil, err
	}

	packet := &models.BillingPacket{UserID: userID, Requests: requests}

	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	_, err = tx.
		NewInsert().
		Model(packet).
		Exec(ctx)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return nil, err
		}
		return nil, err
	}

	billing.Requests += requests
	_, err = tx.
		NewUpdate().
		Model(billing).
		WherePK().
		Column("requests").
		Returning("*").
		Exec(ctx)

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return packet, nil
}

func (s *storeImpl) BillingSmartRegisterRequest(ctx context.Context, userID int64, query string, risk int, category string, network string) (*models.Billing, error) {
	billing, err := s.BillingGetOrCreate(ctx, userID)
	if err != nil {
		return nil, err
	}

	var count int
	err = s.db.
		NewSelect().
		ColumnExpr("count(id)").
		Model((*models.BillingRequest)(nil)).
		Where("query = ? AND user_id = ? AND network = ?", query, userID, network).
		Scan(ctx, &count)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	if count > 0 {
		return billing, nil
	}

	billingRequest := &models.BillingRequest{
		UserID:   userID,
		Query:    query,
		Risk:     risk,
		Category: category,
		Network:  network,
	}

	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	_, err = tx.
		NewInsert().
		Model(billingRequest).
		Exec(ctx)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return nil, err
		}
		return nil, err
	}

	billing.Requests -= 1
	_, err = tx.
		NewUpdate().
		Model(billing).
		WherePK().
		Column("requests").
		Returning("*").
		Exec(ctx)

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return billing, nil
}

func (s *storeImpl) BillingRegisterRequest(ctx context.Context, userID int64, query string, risk int, category string, network string) (*models.Billing, error) {
	billing, err := s.BillingGetOrCreate(ctx, userID)
	if err != nil {
		return nil, err
	}

	if billing.Requests == 0 {
		return nil, ErrInsufficientBalance
	}

	billingRequest := &models.BillingRequest{
		UserID:   userID,
		Query:    query,
		Risk:     risk,
		Category: category,
		Network:  network,
	}

	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	_, err = tx.
		NewInsert().
		Model(billingRequest).
		Exec(ctx)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return nil, err
		}
		return nil, err
	}

	billing.Requests -= 1
	_, err = tx.
		NewUpdate().
		Model(billing).
		WherePK().
		Column("requests").
		Returning("*").
		Exec(ctx)

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return billing, nil
}
