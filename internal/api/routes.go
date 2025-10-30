package api

import (
	"github.com/joaoparis/backendcsbingo/internal/api/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/player/random", handlers.GetRandomPlayer)
	api.Get("/board/random", handlers.GetRandomBoard)
}
