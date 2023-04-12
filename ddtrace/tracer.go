package ddtrace

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

type tracer struct {
	base                   trace.Tracer
	operationName          string
	operationNameFormatter OperationNameFormatter
	spanNameFormatter      SpanNameFormatter
}

func NewTracer(base trace.Tracer, operationName string, opts ...TracerOption) trace.Tracer {
	ret := &tracer{
		base:                   base,
		operationName:          operationName,
		operationNameFormatter: defaultOperationNameFormatter,
		spanNameFormatter:      defaultSpanNameFormatter,
	}
	for _, opt := range opts {
		opt(ret)
	}
	return ret
}

func (d *tracer) Start(ctx context.Context, spanName string,
	opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	operationName := d.operationNameFormatter(ctx, d.operationName)
	spanName = d.spanNameFormatter(ctx, operationName, spanName)
	rctx, span := d.base.Start(ctx, spanName, opts...)
	span.SetAttributes(
		DDOperationNameKey.String(operationName),
		DDResourceNameKey.String(spanName),
	)
	return rctx, span
}

type TracerOption func(*tracer)

func WithTracerOperationNameFormatter(operationNameFormatter OperationNameFormatter) TracerOption {
	return func(p *tracer) {
		p.operationNameFormatter = operationNameFormatter
	}
}

func WithTracerSpanNameFormatter(spanNameFormatter SpanNameFormatter) TracerOption {
	return func(p *tracer) {
		p.spanNameFormatter = spanNameFormatter
	}
}
