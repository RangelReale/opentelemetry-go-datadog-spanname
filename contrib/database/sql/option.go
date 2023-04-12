package sql

import "go.opentelemetry.io/otel/trace"

type options struct {
	tracerProvider trace.TracerProvider
	operationName  string
}

type Option func(*options)

func WithTracerProvider(tracerProvider trace.TracerProvider) Option {
	return func(o *options) {
		o.tracerProvider = tracerProvider
	}
}

func WithOperationName(name string) Option {
	return func(o *options) {
		o.operationName = name
	}
}
