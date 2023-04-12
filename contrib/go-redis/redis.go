// Package redis should be used with:
// github.com/redis/go-redis
package redis

import (
	"context"
	"strings"

	"github.com/RangelReale/opentelemetry-go-datadog-spanname/ddtrace"
	"go.opentelemetry.io/otel/trace"
)

func New(opts ...Option) trace.TracerProvider {
	optns := options{
		operationName: "redis.command",
	}
	for _, opt := range opts {
		opt(&optns)
	}
	return ddtrace.NewTracerProvider(optns.operationName,
		ddtrace.WithTracerProvider(optns.tracerProvider),
		ddtrace.WithSpanNameFormatter(func(ctx context.Context, operationName string, spanName string) string {
			// add a "redis." prefix to span names when not available
			if strings.HasPrefix(spanName, "redis.") {
				return spanName
			}
			return "redis." + spanName
		}))
}
