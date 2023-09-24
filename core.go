package bamboo

import (
	"context"

	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("github.com/kujilabo/bamboo-root")

type BambooRequestProducer interface {
	Produce(ctx context.Context, resultChannel string, heartbeatIntervalSec int, jobTimeoutSec int, headers map[string]string, data []byte) error
	Close(ctx context.Context) error
}

type BambooResultSubscriber interface {
	Ping(ctx context.Context) error
	Subscribe(ctx context.Context, resultChannel string, jobTimeoutSec int) ([]byte, error)
}

type BambooWorker interface {
	Run(ctx context.Context) error
}

type WorkerFunc func(ctx context.Context, headers map[string]string, data []byte, aborted <-chan interface{}) ([]byte, error)

type LogConfigFunc func(ctx context.Context, headers map[string]string) context.Context
