package otlp

import (
	"context"
	"gitlab.com/rubin-dev/api/pkg/models"
	"gitlab.com/rubin-dev/api/pkg/service"
)

var _ service.AccessRequestService = (*metricService)(nil)

func (m *metricService) AccessRequest(ctx context.Context, user *models.User) error {
	ctx, t := m.tracer.Start(ctx, "service.AccessRequest")
	defer t.End()

	return m.s.AccessRequest(ctx, user)
}
