package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ysnbyzli/go-redis-example/handlers"
)

func TodoRoutes(app *fiber.App) {

	routes := app.Group("/todos")

	routes.Get("/", handlers.GetAll)

	routes.Post("/", handlers.Create)

	routes.Put("/", handlers.Update)

	routes.Delete("/", handlers.Delete)

}
