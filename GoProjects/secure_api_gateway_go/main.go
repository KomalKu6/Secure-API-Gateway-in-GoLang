package main

import (
	"github.com/gofiber/fiber/v2"
	"secure_api_gateway_go/handlers"
	"secure_api_gateway_go/middleware"
)

func main() {
	app := fiber.New()

	app.Post("/login", handlers.Login)
	app.Get("/user", middleware.JWTMiddleware("user"), handlers.UserPage)
	app.Get("/admin", middleware.JWTMiddleware("admin"), handlers.AdminPage)

	app.Listen(":3000")
}
