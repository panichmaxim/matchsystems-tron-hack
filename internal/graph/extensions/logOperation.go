package extensions

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/trace"
)

func NewOperationLoggerExtension() OperationLoggerExtension {
	return OperationLoggerExtension{}
}

type OperationLoggerExtension struct {
	tracer trace.Tracer
}

var _ interface {
	graphql.HandlerExtension
	graphql.OperationInterceptor
} = OperationLoggerExtension{}

func (a OperationLoggerExtension) ExtensionName() string {
	return "OperationTracing"
}

func (a OperationLoggerExtension) Validate(schema graphql.ExecutableSchema) error {
	return nil
}

func (a OperationLoggerExtension) InterceptOperation(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
	operationContext := graphql.GetOperationContext(ctx)

	log.Info().Str("operation", operationContext.OperationName).Msgf("graphql operation")

	return next(ctx)
}
