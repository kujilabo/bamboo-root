package bamboo

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/proto"
)

type ByteArreayResult struct {
	Value []byte
	Error error
}

type RedisResultSubscriber struct {
	workerName string
	subscriber redis.UniversalClient
}

func NewRedisResultSubscriber(ctx context.Context, workerName string, subscriberConfig *redis.UniversalOptions) BambooResultSubscriber {
	subscriber := redis.NewUniversalClient(subscriberConfig)

	return &RedisResultSubscriber{
		workerName: workerName,
		subscriber: subscriber,
	}
}

func (s *RedisResultSubscriber) Ping(ctx context.Context) error {
	if _, err := s.subscriber.Ping(ctx).Result(); err != nil {
		return err
	}

	return nil
}

func (s *RedisResultSubscriber) Subscribe(ctx context.Context, resultChannel string, timeout time.Duration) ([]byte, error) {
	pubsub := s.subscriber.Subscribe(ctx, resultChannel)
	defer pubsub.Close()
	c1 := make(chan ByteArreayResult, 1)
	done := make(chan interface{})
	defer close(done)

	go func() {
		defer close(c1)
		for {
			select {
			case <-done:
				return
			default:
				msg, err := pubsub.ReceiveMessage(ctx)
				if err != nil {
					c1 <- ByteArreayResult{Value: nil, Error: err}
					return
				}

				respBytes, err := base64.StdEncoding.DecodeString(msg.Payload)
				if err != nil {
					c1 <- ByteArreayResult{Value: nil, Error: err}
					return
				}

				resp := WorkerResponse{}
				if err := proto.Unmarshal(respBytes, &resp); err != nil {
					c1 <- ByteArreayResult{Value: nil, Error: err}

					return
				}
				c1 <- ByteArreayResult{Value: resp.Data, Error: nil}
			}
		}
	}()

	go func() {
		pulseInterval := time.Second * 3
		pulse := time.Tick(pulseInterval)
		for {
			select {
			case <-done:
				fmt.Println("CLOSE")
				return
			case <-pulse:
				fmt.Println("HEARTBEAT")
			}
		}
	}()

	select {
	case res := <-c1:
		if res.Error != nil {
			return nil, res.Error
		}
		return res.Value, nil
	case <-time.After(timeout):
		return nil, errors.New("timeout")
	}
}
