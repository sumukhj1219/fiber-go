package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sumukhj1219/fiber-go/config"
	"github.com/sumukhj1219/fiber-go/routes"
)

func main() {
	config.ConnectDB()
	defer config.DB.Close()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Content-Type, Authorization",
		AllowCredentials: true,
	}))

	routes.AuthRoutes(app)
	routes.MonitorRoutes(app)
	routes.TrackerRoutes(app)

	app.Listen(":5000")
}
