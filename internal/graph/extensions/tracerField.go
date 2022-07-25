package extensions

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel/attribute"

	"github.com/99designs/gqlgen/graphql"
	"go.opentelemetry.io/otel/trace"
)

func NewFieldTracerExtension(tracer trace.Tracer) FieldTracer {
	return FieldTracer{tracer: tracer}
}

type FieldTracer struct {
	tracer trace.Tracer
}

var _ interface {
	graphql.HandlerExtension
	graphql.FieldInterceptor
} = FieldTracer{}

func (a FieldTracer) ExtensionName() string {
	return "OpenTracing"
}

func (a FieldTracer) Validate(schema graphql.ExecutableSchema) error {
	return nil
}

func (a FieldTracer) InterceptField(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
	fieldContext := graphql.GetFieldContext(ctx)

	ctx, t := a.tracer.Start(ctx, fmt.Sprintf("graphql %s.%s", fieldContext.Object, fieldContext.Field.Name))
	t.SetAttributes(
		attribute.KeyValue{Key: "resolver.object", Value: attribute.StringValue(fieldContext.Object)},
		attribute.KeyValue{Key: "resolver.field", Value: attribute.StringValue(fieldContext.Field.Name)},
	)
	defer t.End()

	return next(ctx)
}
