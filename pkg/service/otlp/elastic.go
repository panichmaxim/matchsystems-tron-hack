package otlp

import (
	"context"
	"gitlab.com/rubin-dev/api/pkg/elastic"
	"gitlab.com/rubin-dev/api/pkg/service"
)

var _ service.ElasticService = (*metricService)(nil)

func (m *metricService) Search(ctx context.Context, value string, page, limit int) ([]*elastic.Entity, int, error) {
	ctx, t := m.tracer.Start(ctx, "service.Search")
	defer t.End()

	return m.s.Search(ctx, value, page, limit)
}

func (m *metricService) SearchCount(ctx context.Context, value string) (int, error) {
	ctx, t := m.tracer.Start(ctx, "service.SearchCount")
	defer t.End()

	return m.s.SearchCount(ctx, value)
}
