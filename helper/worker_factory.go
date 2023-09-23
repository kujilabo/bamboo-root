package helper

import (
	"errors"

	"github.com/redis/go-redis/v9"

	"github.com/kujilabo/bamboo-root"
)

func CreateBambooWorker(cfg *WorkerConfig, workerFunc bamboo.WorkerFunc) (bamboo.BambooWorker, error) {
	if cfg.Consumer.Type == "redis" && cfg.Publisher.Type == "redis" {
		return bamboo.NewRedisRedisBambooWorker(redis.UniversalOptions{
			Addrs:    cfg.Consumer.Redis.Addrs,
			Password: cfg.Consumer.Redis.Password,
		}, cfg.Consumer.Redis.Channel, redis.UniversalOptions{
			Addrs:    cfg.Publisher.Redis.Addrs,
			Password: cfg.Publisher.Redis.Password,
		}, workerFunc, cfg.NumWorkers), nil
	}

	return nil, errors.New("Invalid")
}
