package main

import (
	"github.com/OrangIPA/ukekehfrozekakhyr/controller"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Post("/login", controller.LoginUser)

}

func RestrictedRoutes(app *fiber.App) {
	// USER
	app.Get("/user", controller.GetAllUser)
	app.Get("/user/:id", controller.GetUserById)
	app.Post("/user", controller.CreateUser)
	app.Put("/user/:id", controller.UpdateUser)
	app.Delete("/user/:id", controller.DeleteUser)

	// MENU
	app.Get("/menu", controller.GetAllMenu)
	app.Get("/menu/:id", controller.GetMenuById)
	app.Post("/menu", controller.CreateMenu)
	app.Put("/menu/:id", controller.UpdateMenu)
	app.Delete("/menu/:id", controller.DeleteMenu)

	// MEJA
	app.Get("/meja", controller.GetAllMeja)
	app.Get("/meja/:id", controller.GetMejaById)
	app.Post("/meja", controller.CreateMeja)
	app.Put("/meja/:id", controller.UpdateMeja)
	app.Delete("/meja/:id", controller.DeleteMeja)
}