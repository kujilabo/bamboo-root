package bamboo

import (
	"context"
	"encoding/base64"

	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/proto"

	"github.com/kujilabo/bamboo-root/internal"
)

type RedisJob interface {
	Run(ctx context.Context) error
}

type redisJob struct {
	carrier          propagation.MapCarrier
	workerFunc       WorkerFunc
	headers          map[string]string
	parameter        []byte
	publisherOptions redis.UniversalOptions
	resultChannel    string
}

func NewRedisJob(ctx context.Context, carrier propagation.MapCarrier, workerFunc WorkerFunc, headers map[string]string, parameter []byte, publisherOptions redis.UniversalOptions, resultChannel string) RedisJob {
	return &redisJob{
		carrier:          carrier,
		publisherOptions: publisherOptions,
		workerFunc:       workerFunc,
		parameter:        parameter,
		resultChannel:    resultChannel,
	}
}

func (j *redisJob) Run(ctx context.Context) error {
	propagator := otel.GetTextMapPropagator()
	ctx = propagator.Extract(ctx, j.carrier)

	attrs := make([]attribute.KeyValue, 0)
	for k, v := range j.headers {
		attrs = append(attrs, attribute.KeyValue{Key: attribute.Key(k), Value: attribute.StringValue(v)})
	}

	opts := []trace.SpanStartOption{
		trace.WithAttributes(attrs...),
		trace.WithSpanKind(trace.SpanKindConsumer),
	}
	ctx, span := tracer.Start(ctx, "Run", opts...)
	defer span.End()

	result, err := j.workerFunc(ctx, j.headers, j.parameter)
	if err != nil {
		return internal.Errorf("workerFunc. err: %w", err)
	}

	resp := WorkerResponse{Data: result}
	respBytes, err := proto.Marshal(&resp)
	if err != nil {
		return internal.Errorf("proto.Marshal. err: %w", err)
	}

	respStr := base64.StdEncoding.EncodeToString(respBytes)

	publisher := redis.NewUniversalClient(&j.publisherOptions)
	defer publisher.Close()

	if _, err := publisher.Publish(ctx, j.resultChannel, respStr).Result(); err != nil {
		return internal.Errorf("publisher.Publish. err: %w", err)
	}

	return nil
}
