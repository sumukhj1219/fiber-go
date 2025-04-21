package controllers

import (
	"context"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sumukhj1219/fiber-go/config"
	"github.com/sumukhj1219/fiber-go/models"
)

func Register(c *fiber.Ctx) error {
	ctx := context.Background()
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	_, err := config.DB.Exec(ctx, "INSERT INTO users (username, email) VALUES ($1, $2)", user.Username, user.Email)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Cannot create user"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "User created successfully"})
}

func Login(c *fiber.Ctx) error {
	ctx := context.Background()
	type Payload struct {
		Email    string `json:"email"`
		Username string `json:"username"`
	}
	var payload Payload

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	var user models.User
	err := config.DB.QueryRow(ctx, "SELECT id, username, email FROM users WHERE email = $1", payload.Email).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":   user.ID,
		"email":    user.Email,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not generate token"})
	}

	cookieAuth := new(fiber.Cookie)
	cookieAuth.Name = "token"
	cookieAuth.Value = tokenString
	cookieAuth.Expires = time.Now().Add(time.Hour * 24)
	cookieAuth.MaxAge = 86400
	cookieAuth.HTTPOnly = true
	cookieAuth.Secure = false
	cookieAuth.SameSite = fiber.CookieSameSiteNoneMode

	c.Cookie(cookieAuth)

	return c.JSON(fiber.Map{"message": "Login successful"})
}

func Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
		Secure:   false,
		SameSite: fiber.CookieSameSiteNoneMode,
	})

	return c.JSON(fiber.Map{"message": "Logged out successfully"})
}
