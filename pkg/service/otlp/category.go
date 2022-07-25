package otlp

import (
	"context"
	"gitlab.com/rubin-dev/api/pkg/models"
	"gitlab.com/rubin-dev/api/pkg/service"
)

var _ service.CategoryService = (*metricService)(nil)

func (m *metricService) CategoryList(ctx context.Context) ([]*models.Category, int, error) {
	ctx, t := m.tracer.Start(ctx, "service.CategoryList")
	defer t.End()

	return m.s.CategoryList(ctx)
}

func (m *metricService) CategoryFindByID(ctx context.Context, id int64) (*models.Category, error) {
	ctx, t := m.tracer.Start(ctx, "service.CategoryFindByID")
	defer t.End()

	return m.s.CategoryFindByID(ctx, id)
}

func (m *metricService) CategoryCreate(ctx context.Context, category *models.Category) error {
	ctx, t := m.tracer.Start(ctx, "service.CategoryCreate")
	defer t.End()

	return m.s.CategoryCreate(ctx, category)
}

func (m *metricService) CategoryUpdate(ctx context.Context, category *models.Category, columns ...string) error {
	ctx, t := m.tracer.Start(ctx, "service.CategoryUpdate")
	defer t.End()

	return m.s.CategoryUpdate(ctx, category, columns...)
}

func (m *metricService) CategoryRemoveByID(ctx context.Context, id int64) error {
	ctx, t := m.tracer.Start(ctx, "service.CategoryRemoveByID")
	defer t.End()

	return m.s.CategoryRemoveByID(ctx, id)
}
