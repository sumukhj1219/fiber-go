package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sumukhj1219/fiber-go/controllers"
)

func AuthRoutes(app *fiber.App) {
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)
	app.Post("/logout", controllers.Logout)
}
