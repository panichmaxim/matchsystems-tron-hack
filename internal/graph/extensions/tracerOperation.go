package extensions

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"go.opentelemetry.io/otel/trace"
)

func NewOperationTracerExtension(tracer trace.Tracer) OperationTracer {
	return OperationTracer{tracer: tracer}
}

type OperationTracer struct {
	tracer trace.Tracer
}

var _ interface {
	graphql.HandlerExtension
	graphql.OperationInterceptor
} = OperationTracer{}

func (a OperationTracer) ExtensionName() string {
	return "OperationTracing"
}

func (a OperationTracer) Validate(schema graphql.ExecutableSchema) error {
	return nil
}

func (a OperationTracer) InterceptOperation(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
	operationContext := graphql.GetOperationContext(ctx)

	ctx, t := a.tracer.Start(ctx, fmt.Sprintf("graphql %s", operationContext.OperationName))
	defer t.End()

	return next(ctx)
}
