package http

import (
	"github.com/RangelReale/opentelemetry-go-datadog-spanname/ddtrace"
	"go.opentelemetry.io/otel/trace"
)

// NewHandler should be used with otelhttp.NewHandler
func NewHandler(opts ...Option) trace.TracerProvider {
	optns := options{
		operationName: "http.request",
	}
	for _, opt := range opts {
		opt(&optns)
	}
	return ddtrace.NewTracerProvider(optns.operationName,
		ddtrace.WithTracerProvider(optns.tracerProvider))
}
