package routes

import (
	"github.com/OrangIPA/ukekehfrozekakhyr/db"
	"github.com/OrangIPA/ukekehfrozekakhyr/helper"
	"github.com/OrangIPA/ukekehfrozekakhyr/model"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/forbidden", func(c *fiber.Ctx) error {
		return c.Render("forbidden", fiber.Map{})
	})
	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("login", fiber.Map{})
	})
}

func RestrictedRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		claims := helper.TokenClaims(c)
		role := claims["role"].(string)
		userId := claims["userId"].(uint)
		var user model.User
		if err := db.DB.First(&user, userId).Error; err != nil {
			return err
		}
		switch role {
		case "kasir":
			c.Render("dashboard/kasir", fiber.Map{
				"username": user.Username,
			})
		}
		return c.Render("dashboard", fiber.Map{})
	})
}