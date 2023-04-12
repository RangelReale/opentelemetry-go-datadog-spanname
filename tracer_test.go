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
		instrumentName        string // default is "ddspanname.instrument"
		operationName         string // default is "ddspanname.operation"
		spanName              string // default is "ddspanname.span"
		expectedOperationName string
		expectedResourceName  string
	}{
		{
			expectedOperationName: "ddspanname.operation",
			expectedResourceName:  "ddspanname.span",
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

		ddtracer := NewTracer(tracer, test.operationName)

		_, span := ddtracer.Start(context.Background(), test.spanName)
		span.End()

		spans := sr.Ended()
		require.Len(t, spans, 1)

		for _, testSpan := range spans {
			switch testSpan.Name() {
			case test.spanName:
				require.Contains(t, testSpan.Attributes(), DDOperationNameKey.String(test.expectedOperationName))
				require.Contains(t, testSpan.Attributes(), DDResourceNameKey.String(test.expectedResourceName))
			default:
				t.Fatalf("unexpected span name %s", testSpan.Name())
			}
		}
	}
}
