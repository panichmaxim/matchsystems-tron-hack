package otlp

import (
	"context"
	"gitlab.com/rubin-dev/api/pkg/models"
	"gitlab.com/rubin-dev/api/pkg/service"
)

var _ service.BillingService = (*metricService)(nil)

func (m *metricService) BillingHistoryRequestsCount(ctx context.Context, userID int64) (int, error) {
	ctx, t := m.tracer.Start(ctx, "service.BillingHistoryRequestsCount")
	defer t.End()

	return m.s.BillingHistoryRequestsCount(ctx, userID)
}

func (m *metricService) BillingKeyRemove(ctx context.Context, userID, id int64) error {
	ctx, t := m.tracer.Start(ctx, "service.BillingKeyRemove")
	defer t.End()

	return m.s.BillingKeyRemove(ctx, userID, id)
}

func (m *metricService) BillingKeyList(ctx context.Context, userID int64, page, pageSize int) ([]*models.BillingKey, int, error) {
	ctx, t := m.tracer.Start(ctx, "service.BillingKeyList")
	defer t.End()

	return m.s.BillingKeyList(ctx, userID, page, pageSize)
}

func (m *metricService) BillingAddRequests(ctx context.Context, userID int64, requests int) (*models.BillingPacket, error) {
	ctx, t := m.tracer.Start(ctx, "service.BillingAddRequests")
	defer t.End()

	return m.s.BillingAddRequests(ctx, userID, requests)
}

func (m *metricService) BillingRegisterRequest(ctx context.Context, userID int64, query string, risk int, category string, network string) (*models.Billing, error) {
	ctx, t := m.tracer.Start(ctx, "service.BillingRegisterRequest")
	defer t.End()

	return m.s.BillingRegisterRequest(ctx, userID, query, risk, category, network)
}

func (m *metricService) BillingSmartRegisterRequest(ctx context.Context, userID int64, query string, risk int, category string, network string) (*models.Billing, error) {
	ctx, t := m.tracer.Start(ctx, "service.BillingSmartRegisterRequest")
	defer t.End()

	return m.s.BillingSmartRegisterRequest(ctx, userID, query, risk, category, network)
}

func (m *metricService) BillingGetOrCreate(ctx context.Context, userID int64) (*models.Billing, error) {
	ctx, t := m.tracer.Start(ctx, "service.BillingGetOrCreate")
	defer t.End()

	return m.s.BillingGetOrCreate(ctx, userID)
}

func (m *metricService) BillingFindApiKey(ctx context.Context, apiKey string) (*models.BillingKey, error) {
	ctx, t := m.tracer.Start(ctx, "service.BillingFindApiKey")
	defer t.End()

	return m.s.BillingFindApiKey(ctx, apiKey)
}

func (m *metricService) BillingCreateApiKey(ctx context.Context, userID int64) (*models.BillingKey, error) {
	ctx, t := m.tracer.Start(ctx, "service.BillingCreateApiKey")
	defer t.End()

	return m.s.BillingCreateApiKey(ctx, userID)
}

func (m *metricService) BillingHistoryList(ctx context.Context, userID int64, page int, pageSize int) ([]*models.BillingRequest, int, error) {
	ctx, t := m.tracer.Start(ctx, "service.BillingHistoryList")
	defer t.End()

	return m.s.BillingHistoryList(ctx, userID, page, pageSize)
}
