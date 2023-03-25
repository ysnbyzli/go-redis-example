package main

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ysnbyzli/go-redis-example/cache"
	"github.com/ysnbyzli/go-redis-example/configs"
	"github.com/ysnbyzli/go-redis-example/middlewares"
	"github.com/ysnbyzli/go-redis-example/routes"

	_ "github.com/ysnbyzli/go-redis-example/swagger"
)

// @title Go & Redis Example API
// @version 1.0
// @description Adding food to basket using go and redis
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
// @securityDefinitions.basic BasicAuth
// @name Authorization
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
	routes.SwaggerRoutes(app)

	app.Listen(":3000")

}
