package cache

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/ysnbyzli/go-redis-example/services"
)

var ctx = context.Background()

func RedisClient() (*services.RedisService, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	ping := rdb.Ping(ctx)

	_, err := ping.Result()

	if err != nil {
		return nil, err
	}

	return &services.RedisService{
		Client: rdb,
	}, nil

}
