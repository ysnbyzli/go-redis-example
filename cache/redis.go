package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	client *redis.Client
}

func RedisClient() (*RedisRepository, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	ping := rdb.Ping(context.Background())

	_, err := ping.Result()

	if err != nil {
		return nil, err
	}

	return &RedisRepository{
		client: rdb,
	}, nil

}

func (repository *RedisRepository) Get(key string) *string {
	val, err := repository.client.Get(context.Background(), key).Result()

	if err != nil {
		return nil
	}

	return &val
}

func (repository *RedisRepository) SetKey(key string, data []byte) error {
	err := repository.client.Set(context.Background(), key, data, 5*time.Minute).Err()
	if err != nil {
		panic(err)
	}
	return nil
}

func (repository *RedisRepository) GetKey(key string) ([]byte, error) {

	status := repository.client.Get(context.Background(), key)

	if status.Err() != nil {
		if status.Err() == redis.Nil {
			return nil, fmt.Errorf("%s: key not existing", key)
		}
		return nil, status.Err()
	}

	res, err := status.Bytes()

	if err != nil {
		return nil, err
	}

	return res, nil

}
