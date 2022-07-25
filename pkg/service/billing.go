package service

import (
	"context"
	"gitlab.com/rubin-dev/api/pkg/models"
)

type BillingService interface {
	BillingAddRequests(ctx context.Context, userID int64, requests int) (*models.BillingPacket, error)
	BillingRegisterRequest(ctx context.Context, userID int64, query string, risk int, category string, network string) (*models.Billing, error)
	BillingGetOrCreate(ctx context.Context, userID int64) (*models.Billing, error)
	BillingFindApiKey(ctx context.Context, apiKey string) (*models.BillingKey, error)
	BillingCreateApiKey(ctx context.Context, userID int64) (*models.BillingKey, error)
	BillingSmartRegisterRequest(ctx context.Context, userID int64, query string, risk int, category string, network string) (*models.Billing, error)
	BillingHistoryList(ctx context.Context, userID int64, page, pageSize int) ([]*models.BillingRequest, int, error)
	BillingKeyList(ctx context.Context, userID int64, page, pageSize int) ([]*models.BillingKey, int, error)
	BillingKeyRemove(ctx context.Context, userID, id int64) error
	BillingHistoryRequestsCount(ctx context.Context, userID int64) (int, error)
}

var _ BillingService = (*serviceImpl)(nil)

func (s *serviceImpl) BillingHistoryRequestsCount(ctx context.Context, userID int64) (int, error) {
	return s.s.BillingHistoryRequestsCount(ctx, userID)
}

func (s *serviceImpl) BillingKeyRemove(ctx context.Context, userID, id int64) error {
	return s.s.BillingKeyRemove(ctx, userID, id)
}

func (s *serviceImpl) BillingKeyList(ctx context.Context, userID int64, page, pageSize int) ([]*models.BillingKey, int, error) {
	limit, offset := buildLimitOffset(page, pageSize)
	return s.s.BillingKeyList(ctx, userID, limit, offset)
}

func (s *serviceImpl) BillingGetOrCreate(ctx context.Context, userID int64) (*models.Billing, error) {
	return s.s.BillingGetOrCreate(ctx, userID)
}

func (s *serviceImpl) BillingAddRequests(ctx context.Context, userID int64, requests int) (*models.BillingPacket, error) {
	return s.s.BillingAddRequests(ctx, userID, requests)
}

func (s *serviceImpl) BillingRegisterRequest(ctx context.Context, userID int64, query string, risk int, category string, network string) (*models.Billing, error) {
	return s.s.BillingRegisterRequest(ctx, userID, query, risk, category, network)
}

func (s *serviceImpl) BillingSmartRegisterRequest(ctx context.Context, userID int64, query string, risk int, category string, network string) (*models.Billing, error) {
	return s.s.BillingSmartRegisterRequest(ctx, userID, query, risk, category, network)
}

func (s *serviceImpl) BillingFindApiKey(ctx context.Context, apiKey string) (*models.BillingKey, error) {
	return s.s.BillingFindApiKey(ctx, apiKey)
}

func (s *serviceImpl) BillingCreateApiKey(ctx context.Context, userID int64) (*models.BillingKey, error) {
	return s.s.BillingCreateApiKey(ctx, userID)
}

func (s *serviceImpl) BillingHistoryList(ctx context.Context, userID int64, page int, pageSize int) ([]*models.BillingRequest, int, error) {
	limit, offset := buildLimitOffset(page, pageSize)
	return s.s.BillingHistoryList(ctx, userID, limit, offset)
}
