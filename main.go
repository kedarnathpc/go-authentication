package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kedarnathpc/go-auth/database"
	"github.com/kedarnathpc/go-auth/routes"
)

func main() {
	database.Connect()
	app := fiber.New()
	routes.Setup(app)
	app.Listen(":8000")
}
