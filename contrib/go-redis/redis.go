package redis

import (
	"context"
	"strings"

	ddspanname "github.com/RangelReale/opentelemetry-go-datadog-spanname"
	"go.opentelemetry.io/otel/trace"
)

// New should be used with [github.com/redis/go-redis/extra/redisotel/v9].
func New(opts ...Option) trace.TracerProvider {
	optns := options{
		operationName: "redis.command",
	}
	for _, opt := range opts {
		opt(&optns)
	}
	return ddspanname.NewTracerProvider(optns.operationName,
		ddspanname.WithTracerProvider(optns.tracerProvider),
		ddspanname.WithSpanNameFormatter(func(ctx context.Context, operationName string, spanName string) string {
			// add a "redis." prefix to span names when not available
			if strings.HasPrefix(spanName, "redis.") {
				return spanName
			}
			return "redis." + spanName
		}))
}
