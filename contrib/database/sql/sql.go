package sql

import (
	ddspanname "github.com/RangelReale/opentelemetry-go-datadog-spanname"
	"go.opentelemetry.io/otel/trace"
)

func New(driverName string, opts ...Option) trace.TracerProvider {
	optns := options{
		operationName: driverName + ".query",
	}
	for _, opt := range opts {
		opt(&optns)
	}
	return ddspanname.NewTracerProvider(optns.operationName,
		ddspanname.WithTracerProvider(optns.tracerProvider),
	)
}
