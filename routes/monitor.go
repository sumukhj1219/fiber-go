package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sumukhj1219/fiber-go/controllers"
)

func MonitorRoutes(app *fiber.App) {
	app.Get("/monitor", controllers.MonitorUptime)
}
