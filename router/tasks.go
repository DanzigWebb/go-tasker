package router

import (
	"github.com/gofiber/fiber/v2"
	"task-app/db"
	"task-app/models"
	"task-app/util"
)

func setupTasksRoutes() {
	TASKS.Use(util.SecureAuth())
	TASKS.Post("/create", handleCreateTask)
}

func handleCreateTask(c *fiber.Ctx) error {
	c.Accepts("application/json")
	c.Accepts("json", "text")

	var t models.TaskApi

	if err := c.BodyParser(&t); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiError{
			Message: "Invalid request data",
		})
	}

	id := c.Locals("id")
	u := new(models.User)
	if res := db.DB.Where("uuid = ?", id).First(&u); res.RowsAffected <= 0 {
		return c.JSON(fiber.Map{"error": true, "general": "Cannot find the User"})
	}

	task := models.Task{
		Title:       t.Title,
		Status:      t.Status,
		Description: t.Description,
		OwnerID:     u.ID,
		OwnerType:   "users",
	}

	result := db.DB.Create(&task).Model(models.Task{})

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiError{
			Message: result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(t)
}
