package ddspanname

import "context"

type OperationNameFormatter func(ctx context.Context, operationName string) string
type SpanNameFormatter func(ctx context.Context, operationName string, spanName string) string

func defaultOperationNameFormatter(_ context.Context, operationName string) string {
	return operationName
}

func defaultSpanNameFormatter(_ context.Context, _ string, spanName string) string {
	return spanName
}
