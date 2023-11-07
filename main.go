package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/kedarnathpc/go-auth/database"
	"github.com/kedarnathpc/go-auth/routes"
)

func main() {
	// Connect to the database
	database.Connect()

	// Create a new Fiber application
	app := fiber.New()

	// Use the CORS middleware to handle cross-origin requests
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	// Set up the routes for your application
	routes.Setup(app)

	// Start the server and listen on port 8000
	app.Listen(":8000")
}
