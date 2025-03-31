package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sumukhj1219/fiber-go/config"
	"github.com/sumukhj1219/fiber-go/models"
	"github.com/sumukhj1219/fiber-go/utils"
)

func Register(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Invalid input"})
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Cannot hash password"})
	}

	_, err = config.DB.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", user.Username, user.Email, hashedPassword)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Cannot create user"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "User created successfully"})
}

func Login(c *fiber.Ctx) error {
	var input models.User
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Input"})
	}

	var user models.User
	err := config.DB.QueryRow("SELECT id, username, password, email FROM users WHERE email = $1", input.Email).Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	if err := utils.CheckPassword(user.Password, input.Password); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	token := jwt
}
