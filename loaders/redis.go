package loaders

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func RedisClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	ping := rdb.Ping(ctx)

	message, err := ping.Result()

	if err != nil {
		log.Fatal(err)
	}

	log.Println(message)

	defer rdb.Close()
}
