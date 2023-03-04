package main

import (
	"github.com/OrangIPA/ukekehfrozekakhyr/controller"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Post("/", controller.NyobakAPI)

	app.Post("/user", controller.CreateUser)

	app.Use(func(c *fiber.Ctx) error {
		c.SendStatus(404) // => 404 "Not Found"
		return nil
	})
}
