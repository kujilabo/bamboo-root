package helper

import "go.opentelemetry.io/otel"

var tracer = otel.Tracer("github.com/kujilabo/bamboo-root/helper")
