// cmd/main.go

package main

import (
	"github.com/diana-yelemes/digital_library/database"
	"github.com/diana-yelemes/digital_library/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize Fiber app and database connection
	database.ConnectDb()
	app := fiber.New()
	routes.SetupRoutes(app)
	// Start the Fiber app
	app.Listen(":3000")
}
