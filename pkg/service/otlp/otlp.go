package otlp

import (
	"context"
	"gitlab.com/rubin-dev/api/pkg/service"
	"go.opentelemetry.io/otel/trace"
)

func NewOTLPService(svc service.Service, tracer trace.Tracer) service.Service {
	return &metricService{svc, tracer}
}

type metricService struct {
	s      service.Service
	tracer trace.Tracer
}

func (m *metricService) Health(ctx context.Context) error {
	ctx, t := m.tracer.Start(ctx, "service.Health")
	defer t.End()

	return m.s.Health(ctx)
}

func (m *metricService) Close(ctx context.Context) error {
	ctx, t := m.tracer.Start(ctx, "service.Close")
	defer t.End()

	return m.s.Close(ctx)
}
