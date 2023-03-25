package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SwaggerRoutes(app *fiber.App) {
	routes := app.Group("/swagger")

	routes.Get("/*", swagger.HandlerDefault)

}
