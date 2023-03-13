package main

import (
	"github.com/OrangIPA/ukekehfrozekakhyr/controller"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Post("/api/login", controller.LoginUser)

}

func RestrictedRoutes(app *fiber.App) {
	// USER
	app.Get("/api/user", controller.GetAllUser)        // Admin only
	app.Get("/api/user/:id", controller.GetUserById)   // Admin only
	app.Post("/api/user", controller.CreateUser)       // Admin only
	app.Put("/api/user/:id", controller.UpdateUser)    // Admin only
	app.Delete("/api/user/:id", controller.DeleteUser) // Admin only

	// MENU
	app.Get("/api/menu", controller.GetAllMenu)        // Any
	app.Get("/api/menu/:id", controller.GetMenuById)   // Any
	app.Post("/api/menu", controller.CreateMenu)       // Admin only
	app.Put("/api/menu/:id", controller.UpdateMenu)    // Admin only
	app.Delete("/api/menu/:id", controller.DeleteMenu) // Admin only

	// MEJA
	app.Get("/api/meja", controller.GetAllMeja)             // Any
	app.Get("/api/meja/:id", controller.GetMejaById)        // Any
	app.Patch("/api/meja/:id", controller.UpdateStatusMeja) // Any
	app.Post("/api/meja", controller.CreateMeja)            // Admin only
	app.Put("/api/meja/:id", controller.UpdateMeja)         // Admin only
	app.Delete("/api/meja/:id", controller.DeleteMeja)      // Admin only

	// TRANSAKSI
	app.Get("/api/transaksi", controller.GetAllTransaksi)             // Any
	app.Get("/api/transaksi/:id", controller.GetTransaksiById)        // Any
	app.Post("/api/transaksi", controller.CreateTransaksi)            // Kasir only
	app.Patch("/api/transaksi/:id", controller.UpdateTransaksiStatus) // Kasir only
}
