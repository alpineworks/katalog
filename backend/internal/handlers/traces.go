package handlers

import "go.opentelemetry.io/otel"

var (
	tracer = otel.Tracer("github.com/alpineworks/katalog/backend/internal/handlers")
)
