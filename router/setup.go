package router

import (
	"github.com/gofiber/fiber/v2"
)

// USER handles all the user routes
var USER fiber.Router

// TASKS handles all the tasks routes
var TASKS fiber.Router

// SetupRoutes setups all the Routes
func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	USER = api.Group("/user")
	setupUserRoutes()

	TASKS = api.Group("/tasks")
	setupTasksRoutes()
}
