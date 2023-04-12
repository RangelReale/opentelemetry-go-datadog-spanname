package sql

import (
	"github.com/RangelReale/opentelemetry-go-datadog-spanname/ddtrace"
	"go.opentelemetry.io/otel/trace"
)

func New(driverName string, opts ...Option) trace.TracerProvider {
	optns := options{
		operationName: driverName + ".query",
	}
	for _, opt := range opts {
		opt(&optns)
	}
	return ddtrace.NewTracerProvider(optns.operationName,
		ddtrace.WithTracerProvider(optns.tracerProvider),
	)
}
