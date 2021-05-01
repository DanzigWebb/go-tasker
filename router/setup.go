package router

import (
	"github.com/gofiber/fiber/v2"
)

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello World!")
}

// SetupRoutes setups all the Routes
func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/", hello)

	SetupRoutesGroup(app, "/api/tasks", setupTaskGroup)
}

func SetupRoutesGroup(app *fiber.App, path string, fn func(app *fiber.App, path string)) {
	fn(app, path)
}
