package handlers

import "github.com/gofiber/fiber/v2"

func GetAll(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func Create(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func Update(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func Delete(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
