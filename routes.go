package main

import (
	"os"

	"github.com/OrangIPA/ukekehfrozekakhyr/controller"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func Routes(app *fiber.App) {
	app.Post("/login", controller.LoginUser)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_KEY")),
	}))

	app.Post("/user", controller.CreateUser)

}
