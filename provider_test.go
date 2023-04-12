package ddspanname

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
)

func TestProvider(t *testing.T) {
	sr := tracetest.NewSpanRecorder()
	tracerProvider := sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(sr))

	ddprovider := NewTracerProvider("ddspanname.operation",
		WithTracerProvider(tracerProvider))
	ddtracer := ddprovider.Tracer("ddspanname.instrument")

	require.IsType(t, &tracer{}, ddtracer)

	_, span := ddtracer.Start(context.Background(), "ddspanname.span")
	span.End()

	spans := sr.Ended()
	require.Len(t, spans, 1)

	require.Equal(t, "ddspanname.span", spans[0].Name())
	require.Contains(t, spans[0].Attributes(), DDOperationNameKey.String("ddspanname.operation"))
	require.Contains(t, spans[0].Attributes(), DDResourceNameKey.String("ddspanname.span"))

}
