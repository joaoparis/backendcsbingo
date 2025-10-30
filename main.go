package main

import (
	"github.com/joaoparis/backendcsbingo/internal/api"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	// 👇 Allow requests from your Flutter web app
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // or "http://localhost:1234" if you want to restrict it
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET,POST,OPTIONS",
	}))

	// register your routes
	api.RegisterRoutes(app)

	app.Listen(":8080")
}
