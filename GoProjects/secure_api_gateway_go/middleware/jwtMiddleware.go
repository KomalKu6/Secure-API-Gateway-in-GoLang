package middleware

import (
	"strings"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("my-secret-key")

func JWTMiddleware(requiredRole string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(401).SendString("Missing or invalid token format")
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			return c.Status(401).SendString("Invalid token")
		}

		claims := token.Claims.(jwt.MapClaims)
		role := claims["role"].(string)

		if role != requiredRole {
			return c.Status(403).SendString("Forbidden: role mismatch")
		}

		return c.Next()
	}
}