package sql

import (
	ddspanname "github.com/RangelReale/opentelemetry-go-datadog-spanname"
	"go.opentelemetry.io/otel/trace"
)

// New should be used with [github.com/XSAM/otelsql], but can be used with any [database/sql] instrument.
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
