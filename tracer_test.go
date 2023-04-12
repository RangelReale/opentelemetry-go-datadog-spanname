package ddspanname

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
)

func TestTracer(t *testing.T) {
	tests := []struct {
		instrumentName string // default is "ddspanname.instrument"
		operationName  string // default is "ddspanname.operation"
		spanName       string // default is "ddspanname.span"
		options        []TracerOption

		expectedSpanName      string
		expectedOperationName string
		expectedResourceName  string
	}{
		{
			expectedSpanName:      "ddspanname.span",
			expectedOperationName: "ddspanname.operation",
			expectedResourceName:  "ddspanname.span",
		},
		{
			options: []TracerOption{
				WithTracerOperationNameFormatter(func(ctx context.Context, operationName string) string {
					return operationName + "Custom"
				}),
			},
			expectedSpanName:      "ddspanname.span",
			expectedOperationName: "ddspanname.operationCustom",
			expectedResourceName:  "ddspanname.span",
		},
		{
			options: []TracerOption{
				WithTracerSpanNameFormatter(func(ctx context.Context, operationName string, spanName string) string {
					return spanName + "Custom"
				}),
			},
			expectedSpanName:      "ddspanname.spanCustom",
			expectedOperationName: "ddspanname.operation",
			expectedResourceName:  "ddspanname.spanCustom",
		},
	}

	for _, test := range tests {
		if test.instrumentName == "" {
			test.instrumentName = "ddspanname.instrument"
		}
		if test.operationName == "" {
			test.operationName = "ddspanname.operation"
		}
		if test.spanName == "" {
			test.spanName = "ddspanname.span"
		}

		sr := tracetest.NewSpanRecorder()
		tracerProvider := sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(sr))
		tracer := tracerProvider.Tracer(test.instrumentName)

		ddtracer := NewTracer(tracer, test.operationName, test.options...)

		_, span := ddtracer.Start(context.Background(), test.spanName)
		span.End()

		spans := sr.Ended()
		require.Len(t, spans, 1)

		require.Equal(t, test.expectedResourceName, spans[0].Name())
		require.Contains(t, spans[0].Attributes(), DDOperationNameKey.String(test.expectedOperationName))
		require.Contains(t, spans[0].Attributes(), DDResourceNameKey.String(test.expectedResourceName))
	}
}
