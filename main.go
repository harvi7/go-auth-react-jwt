package main

import (
	"go-auth-react-jwt/database"
	"go-auth-react-jwt/routes"
	"github.com/gofiber/fiber/v2" 
)
  
func main() {
	database.Connect()

    app := fiber.New()
 
	routes.Setup(app)

	app.Listen(":8000")
}