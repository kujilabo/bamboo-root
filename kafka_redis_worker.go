package bamboo

import (
	"context"
	"errors"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/protobuf/proto"

	"github.com/kujilabo/bamboo-root/internal"
)

type kafkaRedisBambooWorker struct {
	consumerOptions  kafka.ReaderConfig
	publisherOptions redis.UniversalOptions
	workerFunc       WorkerFunc
	numWorkers       int
	propagator       propagation.TextMapPropagator
}

func NewKafkaRedisBambooWorker(consumerOptions kafka.ReaderConfig, publisherOptions redis.UniversalOptions, workerFunc WorkerFunc, numWorkers int) BambooWorker {
	return &kafkaRedisBambooWorker{
		consumerOptions:  consumerOptions,
		publisherOptions: publisherOptions,
		workerFunc:       workerFunc,
		numWorkers:       numWorkers,
		propagator:       otel.GetTextMapPropagator(),
	}
}

func (w *kafkaRedisBambooWorker) ping(ctx context.Context) error {
	if len(w.consumerOptions.Brokers) == 0 {
		return errors.New("broker size is 0")
	}

	conn, err := kafka.Dial("tcp", w.consumerOptions.Brokers[0])
	if err != nil {
		return internal.Errorf("kafka.Dial. err: %w", err)
	}
	defer conn.Close()

	if _, err := conn.ReadPartitions(); err != nil {
		return internal.Errorf("conn.ReadPartitions. err: %w", err)
	}

	publisher := redis.NewUniversalClient(&w.publisherOptions)
	defer publisher.Close()
	if _, err := publisher.Ping(ctx).Result(); err != nil {
		internal.Errorf("publisher.Ping. err: %w", err)
	}

	return nil
}

func (w *kafkaRedisBambooWorker) Run(ctx context.Context) error {
	logger := internal.FromContext(ctx)
	operation := func() error {
		if err := w.ping(ctx); err != nil {
			return err
		}

		dispatcher := internal.NewDispatcher()
		defer dispatcher.Stop(ctx)
		dispatcher.Start(ctx, w.numWorkers)

		r := kafka.NewReader(w.consumerOptions)
		defer r.Close()

		for {
			m, err := r.ReadMessage(ctx)
			if err != nil {
				return internal.Errorf("kafka.ReadMessage. err: %w", err)
			}

			if len(m.Key) == 0 && len(m.Value) == 0 {
				return errors.New("kafka conn error")
			}

			// fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))

			req := WorkerParameter{}
			if err := proto.Unmarshal(m.Value, &req); err != nil {
				logger.Warnf("invalid parameter. failed to proto.Unmarshal. err: %w", err)
				continue
			}

			var carrier propagation.MapCarrier = req.Carrier
			dispatcher.AddJob(NewRedisJob(ctx, carrier, w.workerFunc, req.Headers, req.Data, w.publisherOptions, req.ResultChannel))
		}
	}

	backOff := backoff.NewExponentialBackOff()
	backOff.MaxElapsedTime = 0

	notify := func(err error, d time.Duration) {
		logger.Errorf("kafka reading error. err: %v", err)
	}

	err := backoff.RetryNotify(operation, backOff, notify)
	if err != nil {
		return err
	}

	return nil
}
