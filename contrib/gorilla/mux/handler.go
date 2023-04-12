package mux

import (
	"net/http"

	ddspanname "github.com/RangelReale/opentelemetry-go-datadog-spanname"
	"go.opentelemetry.io/otel/trace"
)

// NewHandler should be used with [github.com/gorilla/mux].
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

// SpanNameFormatter should be used with [otelmux.WithSpanNameFormatter] because
// the default formatter only outputs the route without the http method.
func SpanNameFormatter(routeName string, r *http.Request) string {
	return r.Method + " " + routeName
}
