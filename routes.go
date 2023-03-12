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
	app.Get("/user", controller.GetAllUser)        // Admin only
	app.Get("/user/:id", controller.GetUserById)   // Admin only
	app.Post("/user", controller.CreateUser)       // Admin only
	app.Put("/user/:id", controller.UpdateUser)    // Admin only
	app.Delete("/user/:id", controller.DeleteUser) // Admin only

	// MENU
	app.Get("/menu", controller.GetAllMenu)        // Any
	app.Get("/menu/:id", controller.GetMenuById)   // Any
	app.Post("/menu", controller.CreateMenu)       // Admin only
	app.Put("/menu/:id", controller.UpdateMenu)    // Admin only
	app.Delete("/menu/:id", controller.DeleteMenu) // Admin only

	// MEJA
	app.Get("/meja", controller.GetAllMeja)             // Any
	app.Get("/meja/:id", controller.GetMejaById)        // Any
	app.Patch("/meja/:id", controller.UpdateStatusMeja) // Any
	app.Post("/meja", controller.CreateMeja)            // Admin only
	app.Put("/meja/:id", controller.UpdateMeja)         // Admin only
	app.Delete("/meja/:id", controller.DeleteMeja)      // Admin only

	// TRANSAKSI
	app.Get("/transaksi", controller.GetAllTransaksi)             // Any
	app.Get("/transaksi/:id", controller.GetTransaksiById)        // Any
	app.Post("/transaksi", controller.CreateTransaksi)            // Kasir only
	app.Patch("/transaksi/:id", controller.UpdateTransaksiStatus) // Kasir only
}
