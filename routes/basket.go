package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ysnbyzli/go-redis-example/cache"
	"github.com/ysnbyzli/go-redis-example/handlers"
)

func BasketRoutes(app *fiber.App, redis *cache.RedisRepository) {
	routes := app.Group("/baskets")

	routes.Get("/", handlers.GetBasket(redis))
	routes.Post("/", handlers.SaveBasket(redis))
}
