package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ysnbyzli/go-redis-example/models"
)

// @Summary	Get Foods
// @Tags		Foods
// @Description Returns the list of foods
// @Produce	json
// @Success	200	{object}	[]models.Food
// @Router	/foods [GET]
// @Security Bearer
func GetFoods() fiber.Handler {
	return func(c *fiber.Ctx) error {

		return c.JSON(fiber.Map{
			"foods": models.Foods,
		})
	}
}
