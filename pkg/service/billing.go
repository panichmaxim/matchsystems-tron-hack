package service

import (
	"context"
	"gitlab.com/rubin-dev/api/pkg/models"
	"gitlab.com/rubin-dev/api/pkg/neo4jstore"
	"time"
)

type BillingService interface {
	BillingRisks(ctx context.Context, id int64) ([]*models.BillingRisk, error)
	BillingStatisticsSummary(
		ctx context.Context,
		userID int64,
		from, to *time.Time,
		last *bool,
	) ([]*models.StatisticsSummary, error)
	BillingAddRequests(ctx context.Context, userID int64, requests int) (*models.BillingPacket, error)
	BillingRegisterRequest(ctx context.Context, userID int64, query string, risk *neo4jstore.Risk, network string) (*models.Billing, error)
	BillingGetOrCreate(ctx context.Context, userID int64) (*models.Billing, error)
	BillingFindApiKey(ctx context.Context, apiKey string) (*models.BillingKey, error)
	BillingCreateApiKey(ctx context.Context, userID int64) (*models.BillingKey, error)
	BillingHistoryList(
		ctx context.Context,
		userID int64,
		page, pageSize int,
		from, to *time.Time,
		last *bool,
	) ([]*models.BillingRequest, int, error)
	BillingKeyList(ctx context.Context, userID int64, page, pageSize int) ([]*models.BillingKey, int, error)
	BillingKeyRemove(ctx context.Context, userID, id int64) error
	BillingHistoryRequestsCount(ctx context.Context, userID int64) (int, error)
	BillingStatisticsRisks(
		ctx context.Context,
		userID int64,
		from, to *time.Time,
		network string,
		last *bool, // @todo
	) ([]*models.BillingRisk, error)
	BillingRiskFindCategoriesAndDirectories(
		ctx context.Context,
		userID int64,
		from, to *time.Time,
		network string,
		last *bool,
	) (map[string]int, error)
	BillingStatisticsRiskRanges(
		ctx context.Context,
		userID int64,
		from, to *time.Time,
		network string,
		last *bool,
	) ([]int, error)
	BillingStatisticsNetwork(
		ctx context.Context,
		userID int64,
		from, to *time.Time,
		network string,
		last *bool,
	) (*models.BillingStatisticsBlockchain, error)
}

var _ BillingService = (*serviceImpl)(nil)

func (s *serviceImpl) BillingStatisticsRisks(
	ctx context.Context,
	userID int64,
	from, to *time.Time,
	network string,
	last *bool,
) ([]*models.BillingRisk, error) {
	return s.s.BillingStatisticsRisks(ctx, userID, from, to, network, last)
}

func (s *serviceImpl) BillingStatisticsSummary(
	ctx context.Context,
	userID int64,
	from, to *time.Time,
	last *bool,
) ([]*models.StatisticsSummary, error) {
	return s.s.BillingStatisticsSummary(ctx, userID, from, to, last)
}

func (s *serviceImpl) BillingStatisticsRiskRanges(
	ctx context.Context,
	userID int64,
	from, to *time.Time,
	network string,
	last *bool,
) ([]int, error) {
	return s.s.BillingStatisticsRiskRanges(ctx, userID, from, to, network, last)
}

func (s *serviceImpl) BillingRiskFindCategoriesAndDirectories(
	ctx context.Context,
	userID int64,
	from, to *time.Time,
	network string,
	last *bool,
) (map[string]int, error) {
	return s.s.BillingRiskFindCategoriesAndDirectories(ctx, userID, from, to, network, last)
}

func (s *serviceImpl) BillingRisks(ctx context.Context, id int64) ([]*models.BillingRisk, error) {
	return s.s.BillingRisks(ctx, id)
}

func (s *serviceImpl) BillingStatisticsNetwork(
	ctx context.Context,
	userID int64,
	from, to *time.Time,
	network string,
	last *bool,
) (*models.BillingStatisticsBlockchain, error) {
	return s.s.BillingStatisticsNetwork(ctx, userID, from, to, network, last)
}

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

func (s *serviceImpl) BillingRegisterRequest(ctx context.Context, userID int64, query string, risk *neo4jstore.Risk, network string) (*models.Billing, error) {
	return s.s.BillingRegisterRequest(ctx, userID, query, risk, network)
}

func (s *serviceImpl) BillingFindApiKey(ctx context.Context, apiKey string) (*models.BillingKey, error) {
	return s.s.BillingFindApiKey(ctx, apiKey)
}

func (s *serviceImpl) BillingCreateApiKey(ctx context.Context, userID int64) (*models.BillingKey, error) {
	return s.s.BillingCreateApiKey(ctx, userID)
}

func (s *serviceImpl) BillingHistoryList(
	ctx context.Context,
	userID int64,
	page, pageSize int,
	from, to *time.Time,
	last *bool,
) ([]*models.BillingRequest, int, error) {
	limit, offset := buildLimitOffset(page, pageSize)
	return s.s.BillingHistoryList(ctx, userID, limit, offset, from, to, last)
}
