package http

import (
	ddspanname "github.com/RangelReale/opentelemetry-go-datadog-spanname"
	"go.opentelemetry.io/otel/trace"
)

// NewHandler should be used with [go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp.NewHandler].
func NewHandler(opts ...Option) trace.TracerProvider {
	optns := options{
		operationName: "http.request",
	}
	for _, opt := range opts {
		opt(&optns)
	}
	return ddspanname.NewTracerProvider(optns.operationName,
		ddspanname.WithTracerProvider(optns.tracerProvider))
}
