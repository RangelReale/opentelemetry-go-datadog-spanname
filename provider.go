package ddspanname

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type tracerProvider struct {
	base                   trace.TracerProvider
	operationName          string
	operationNameFormatter OperationNameFormatter
	spanNameFormatter      SpanNameFormatter
}

// NewTracerProvider creates a [trace.TracerProvider] that returns a [trace.Tracer] which
// sets the attributes "operation.name" and "resource.name", so that DataDog span names are formatted
// correctly.
func NewTracerProvider(operationName string, opts ...Option) trace.TracerProvider {
	ret := &tracerProvider{
		base:                   otel.GetTracerProvider(),
		operationName:          operationName,
		operationNameFormatter: defaultOperationNameFormatter,
		spanNameFormatter:      defaultSpanNameFormatter,
	}
	for _, opt := range opts {
		opt(ret)
	}
	return ret
}

func (d *tracerProvider) Tracer(name string, options ...trace.TracerOption) trace.Tracer {
	return NewTracer(d.base.Tracer(name, options...), d.operationName,
		WithTracerOperationNameFormatter(d.operationNameFormatter),
		WithTracerSpanNameFormatter(d.spanNameFormatter))
}

type Option func(*tracerProvider)

// WithTracerProvider sets a base [trace.TracerProvider]. If not set, the global trace provider will be used.
func WithTracerProvider(base trace.TracerProvider) Option {
	return func(p *tracerProvider) {
		if base == nil {
			return
		}
		p.base = base
	}
}

// WithOperationNameFormatter sets a function to customize the operation name.
func WithOperationNameFormatter(operationNameFormatter OperationNameFormatter) Option {
	return func(p *tracerProvider) {
		p.operationNameFormatter = operationNameFormatter
	}
}

// WithSpanNameFormatter sets a function to customize the span name.
func WithSpanNameFormatter(spanNameFormatter SpanNameFormatter) Option {
	return func(p *tracerProvider) {
		p.spanNameFormatter = spanNameFormatter
	}
}
