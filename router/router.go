package router

import (
	loginRoutes "github.com/devtosun/datamorph-go/router/authRoutes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	// Setup the Node Routes
	loginRoutes.SetupNoteRoutes(api)
}
