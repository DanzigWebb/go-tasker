package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"task-app/models"
)

func CreateServer() *fiber.App {
	app := fiber.New()

	return app
}

func main() {
	app := CreateServer()
	app.Use(cors.New())

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/tasks/create", handleCreateTask)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(":3000"))
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
