package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kedarnathpc/go-auth/controllers"
)

// Setup initializes the application routes.
func Setup(app *fiber.App) {
	// Define routes and map them to corresponding controller functions.
	app.Post("/api/register", controllers.Register) // User registration route.
	app.Post("/api/login", controllers.Login)       // User login route.
	app.Get("/api/user", controllers.User)          // User profile route.
	app.Post("/api/logout", controllers.Logout)     // User logout route.
}
