package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ysnbyzli/go-redis-example/handlers"
)

func FoodRoutes(app *fiber.App) {
	routes := app.Group("/foods")

	routes.Get("/", handlers.GetFoods())

}
