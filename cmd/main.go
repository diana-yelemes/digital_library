// cmd/main.go

package main

import (
	"github.com/diana-yelemes/digital_library/database"
	"github.com/diana-yelemes/digital_library/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Initialize Fiber app and database connection
	database.ConnectDb()
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	routes.SetupRoutes(app)

	app.Static("/", "./public")
	// Start the Fiber app
	app.Listen(":3000")
}
