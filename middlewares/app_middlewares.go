package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func AppMiddlewares(app *fiber.App) {
	app.Use(
		cors.New(),
		logger.New(),
		basicauth.New(basicauth.Config{
			Users: map[string]string{
				"yasin": "123456",
			},
		}),
	)
}
