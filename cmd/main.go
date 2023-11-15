// cmd/main.go

package main

import (
	"github.com/diana-yelemes/digital_library/database"
	"github.com/gofiber/fiber"
)

func main() {
	// Initialize Fiber app and database connection
	database.ConnectDb()
	app := fiber.New()
	setupRoutes(app)
	// Start the Fiber app
	app.Listen(3000)
}
