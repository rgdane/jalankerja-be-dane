package routes

import (
	"jk-api/internal/container"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App, c *container.AppContainer) {
	api := app.Group("/api/v1")
	ProjectRoutes(api, c)
	SquadRoutes(api, c)
	UserRoutes(api, c)
}
