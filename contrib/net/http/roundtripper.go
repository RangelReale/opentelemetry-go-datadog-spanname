package http

import (
	"github.com/RangelReale/opentelemetry-go-datadog-spanname/ddtrace"
	"go.opentelemetry.io/otel/trace"
)

func NewRoundTripper(opts ...Option) trace.TracerProvider {
	optns := options{
		operationName: "http.client.request",
	}
	for _, opt := range opts {
		opt(&optns)
	}
	return ddtrace.NewTracerProvider(optns.operationName,
		ddtrace.WithTracerProvider(optns.tracerProvider))
}
