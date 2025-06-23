package handlers

import (
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("my-secret-key")

type LoginRequest struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}

func Login(c *fiber.Ctx) error {
	var data LoginRequest
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).SendString("Invalid input")
	}

	claims := jwt.MapClaims{
		"username": data.Username,
		"role":     data.Role,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(jwtKey)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(fiber.Map{"token": t})
}

func UserPage(c *fiber.Ctx) error {
	return c.SendString("Welcome, user!")
}

func AdminPage(c *fiber.Ctx) error {
	return c.SendString("Welcome, admin!")
}
