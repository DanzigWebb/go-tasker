package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"task-app/db"
	"task-app/router"
)

func CreateServer() *fiber.App {
	app := fiber.New()

	return app
}

func main() {
	db.ConnectToDB()

	app := CreateServer()
	app.Use(cors.New())
	router.SetupRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(":3000"))
}
