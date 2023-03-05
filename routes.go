package main

import (
	"os"

	"github.com/OrangIPA/ukekehfrozekakhyr/controller"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func Routes(app *fiber.App) {
	app.Post("/login", controller.LoginUser)

	// JWT Middleware: Put every route that require JWT AFTER this
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_KEY")),
	}))

	// USER
	app.Get("/user", controller.GetAllUser)
	app.Get("/user/:id", controller.GetUserById)
	app.Post("/user", controller.CreateUser)
	app.Put("/user", controller.UpdateUser)

	// MENU
	app.Post("/menu", controller.CreateMenu)
}
