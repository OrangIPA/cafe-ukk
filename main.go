package main

import (
	"os"

	"github.com/OrangIPA/ukekehfrozekakhyr/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func init() {
	db.LoadEnvVariables()
	db.ConnectToDatabase()
	db.SyncDB()
}

func main() {
	// Load templates
	engine := html.New("./view", ".html")

	// Setup templating engine
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Configure app
	app.Static("/", "./public")

	// Routes
	Routes(app)

	// Start app
	app.Listen(":" + os.Getenv("PORT"))
}
