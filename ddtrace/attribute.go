package ddtrace

import "go.opentelemetry.io/otel/attribute"

const (
	DDOperationNameKey = attribute.Key("operation.name")
	DDResourceNameKey  = attribute.Key("resource.name")
)
