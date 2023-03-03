package controller

import "github.com/gofiber/fiber/v2"

func PostIndex(c *fiber.Ctx) error {
	return c.Render("post/index", fiber.Map{})
}