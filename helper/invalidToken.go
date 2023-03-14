package helper

import "github.com/gofiber/fiber/v2"

func InvalidToken(c *fiber.Ctx, err error) error {
	return c.Redirect("/login")
}