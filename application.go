package main

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ysnbyzli/go-redis-example/cache"
	"github.com/ysnbyzli/go-redis-example/configs"
	"github.com/ysnbyzli/go-redis-example/middlewares"
	"github.com/ysnbyzli/go-redis-example/routes"
)

func main() {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	if err := configs.LoadENV(); err != nil {
		log.Fatalf("Error env: %s", err)
	}

	middlewares.AppMiddlewares(app)

	redis, err := cache.RedisClient()

	if err != nil {
		log.Fatalf("Error redis: %s", err)
	}

	routes.FoodRoutes(app)
	routes.BasketRoutes(app, redis)

	app.Listen(":3000")

}
