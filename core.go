package bamboo

import (
	"context"
	"time"

	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("github.com/kujilabo/bamboo-root")

type BambooRequestProducer interface {
	Produce(ctx context.Context, resultChannel string, headers map[string]string, data []byte) error
	Close(ctx context.Context) error
}

type BambooResultSubscriber interface {
	Ping(ctx context.Context) error
	Subscribe(ctx context.Context, resultChannel string, timeout time.Duration) ([]byte, error)
}

type BambooWorker interface {
	Run(ctx context.Context) error
}

type WorkerFunc func(ctx context.Context, headers map[string]string, data []byte) ([]byte, error)
