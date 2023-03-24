package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ysnbyzli/go-redis-example/configs"
	"github.com/ysnbyzli/go-redis-example/middlewares"
)

func main() {
	app := fiber.New()

	if err := configs.LoadENV(); err != nil {
		log.Fatalf("Error env: %s", err)
	}

	middlewares.AppMiddlewares(app)

	app.Listen(":3000")

}
