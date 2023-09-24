package bamboo

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/kujilabo/bamboo-root/internal"
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

func (s *RedisResultSubscriber) Subscribe(ctx context.Context, resultChannel string, jobTimeoutSec int) ([]byte, error) {
	logger := internal.FromContext(ctx)

	pubsub := s.subscriber.Subscribe(ctx, resultChannel)
	defer pubsub.Close()
	c1 := make(chan ByteArreayResult, 1)
	done := make(chan interface{})
	heartbeat := make(chan int64)

	aborted := make(chan interface{})

	if jobTimeoutSec != 0 {
		time.AfterFunc(time.Duration(jobTimeoutSec)*time.Second, func() {
			logger.Debugf("job timed out. resultChannel: %s", resultChannel)
			close(aborted)
		})
	}

	go func() {
		defer close(c1)
		defer close(done)
		for {
			select {
			case <-done:
				fmt.Println("DONE")
				return
			default:
				fmt.Println("WAIT")
				msg, err := pubsub.ReceiveMessage(ctx)
				if err != nil {
					c1 <- ByteArreayResult{Value: nil, Error: err}
					return
				}
				fmt.Println("RECV")

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
				if resp.Type == ResponseType_HEARTBEAT {
					heartbeat <- time.Now().Unix()
				} else {
					c1 <- ByteArreayResult{Value: resp.Data, Error: nil}
				}
			}
		}
	}()

	go func() {
		var prev int64 = time.Now().Unix()
		for {
			select {
			case <-done:
				fmt.Println("CLOSE")
				return
			case next := <-heartbeat:
				fmt.Println(next - prev)
				fmt.Println("HEARTBEAT<-")
				// done <- struct{}{}
			case <-aborted:
				fmt.Println("CLOSE")
				return
			}
		}
	}()

	select {
	case res := <-c1:
		if res.Error != nil {
			return nil, res.Error
		}
		return res.Value, nil
	case <-aborted:
		return nil, errors.New("timeout")
	}
}
