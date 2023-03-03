package main

import (
	"github.com/OrangIPA/ukekehfrozekakhyr/controller"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/", controller.PostIndex)
}
