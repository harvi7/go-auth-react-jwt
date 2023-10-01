package main

import (
	"github.com/gofiber/fiber/v2"
	"go-auth-react-jwt/database"
	"go-auth-react-jwt/routes"
)

func main() {
	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8000")
}
