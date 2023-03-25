package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ysnbyzli/go-redis-example/cache"
	"github.com/ysnbyzli/go-redis-example/models"
)

func GetFoods() fiber.Handler {
	return func(c *fiber.Ctx) error {

		return c.JSON(fiber.Map{
			"foods": models.Foods,
		})
	}
}

func SaveBasket(redis *cache.RedisRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		basket := new(models.Basket)

		if err := c.BodyParser(&basket); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		user := c.Context().UserValue("username").(string)

		json, err := json.Marshal(basket)

		if err != nil {
			return c.JSON(fiber.Map{
				"status": fiber.ErrInternalServerError,
			})
		}

		rdsErr := redis.SetKey(fmt.Sprintf("basket_%s", user), json)

		if rdsErr != nil {
			return c.JSON(fiber.Map{
				"status": fiber.ErrInternalServerError,
			})
		}

		return c.JSON(fiber.Map{
			"status": fiber.StatusOK,
		})
	}
}

func GetBasket(redis *cache.RedisRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		basket := models.Basket{}
		cache, err := redis.GetKey("basket_yasin")

		if err != nil {
			fmt.Println(err)
		}

		if cache == nil {
			return c.JSON(fiber.Map{
				"foods":       make([]models.Food, 0),
				"total_price": 0,
			})
		}

		if err := json.Unmarshal(cache, &basket); err != nil {
			fmt.Println(err)
			c.Status(fiber.ErrInternalServerError.Code)
			return c.JSON(fiber.Map{
				"foods": fiber.ErrInternalServerError.Message,
			})
		}

		price := 0.0

		for _, food := range basket.Foods {
			price += food.Price * float64(food.Count)
		}

		return c.JSON(fiber.Map{
			"foods":       basket.Foods,
			"total_price": price,
		})
	}
}
