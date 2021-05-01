package router

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"task-app/models"
)

func setupTaskGroup(app *fiber.App, path string) {
	api := app.Group(path)
	api.Post("/create", handleCreateTask)
}

func handleCreateTask(c *fiber.Ctx) error {
	c.Accepts("application/json")
	c.Accepts("json", "text")

	body := c.Body()
	var t models.Task
	err := json.Unmarshal(body, &t)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendString(string(body))
}
