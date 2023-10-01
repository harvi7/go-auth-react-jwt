package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-auth-react-jwt/controllers"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Get("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)
}
