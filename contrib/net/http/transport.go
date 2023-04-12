package http

import (
	"net/http"

	ddspanname "github.com/RangelReale/opentelemetry-go-datadog-spanname"
	"go.opentelemetry.io/otel/trace"
)

// NewTransport should be used with otelhttp.NewTransport.
// Also use TransportSpanNameFormatter with otelhttp.WithSpanNameFormatter for a better span name.
func NewTransport(opts ...Option) trace.TracerProvider {
	optns := options{
		operationName: "http.client.request",
	}
	for _, opt := range opts {
		opt(&optns)
	}
	return ddspanname.NewTracerProvider(optns.operationName,
		ddspanname.WithTracerProvider(optns.tracerProvider))
}

// TransportSpanNameFormatter should be used with otelhttp.WithSpanNameFormatter because
// the default formatter only outputs "HTTP <method>"
func TransportSpanNameFormatter(_ string, r *http.Request) string {
	return r.Method + " " + r.URL.Path
}
