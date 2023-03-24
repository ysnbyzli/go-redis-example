package services

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type RedisService struct {
	*redis.Client
}

func (q *RedisService) Get(key string) (string, error) {
	val, err := q.Client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, err
}

func (q *RedisService) Set(key string, value string) error {
	err := q.Client.Set(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}
