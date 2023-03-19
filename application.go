package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ysnbyzli/go-redis-example/middlewares"
	"github.com/ysnbyzli/go-redis-example/routes"
)

func main() {
	app := fiber.New()

	middlewares.AppMiddlewares(app)

	routes.TodoRoutes(app)

	app.Listen(":3000")

}
