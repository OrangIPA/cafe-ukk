package main

import (
	"os"

	"github.com/OrangIPA/ukekehfrozekakhyr/db"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
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

	// JWT Middleware: Put every route that require JWT AFTER this
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_KEY")),
	}))

	// Restricted routes: routes that need JWT Token
	RestrictedRoutes(app)

	// Not Found
	app.Use(func(c *fiber.Ctx) error {
		c.SendStatus(404) // => 404 "Not Found"
		return nil
	})

	// Start app
	app.Listen(":" + os.Getenv("PORT"))
}
