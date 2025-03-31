package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New(fiber.Config{
		AppName: "Fiber app",
	})

	app.Listen(":3000")
}
