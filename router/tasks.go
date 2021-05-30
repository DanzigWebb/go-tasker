package router

import (
	"github.com/gofiber/fiber/v2"
	"task-app/db"
	"task-app/models"
	"task-app/util"
)

func setupTasksRoutes() {
	TASKS.Use(util.SecureAuth())
	TASKS.Post("/", handleCreateTask)
	TASKS.Patch("/", handleUpdateTask)
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

	u, err := util.GetUserByLocal(c)
	if err != nil {
		return c.JSON(models.DefaultError("Cannot find the User"))
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

	t.ID = task.ID
	t.CreatedAt = task.CreatedAt.String()
	t.UpdatedAt = task.UpdatedAt

	return c.Status(fiber.StatusOK).JSON(t)
}

func handleUpdateTask(c *fiber.Ctx) error {
	c.Accepts("application/json")
	c.Accepts("json", "text")

	var sendError = func(m string, s int) error {
		return models.DefaultError(m).SendStatus(c, s)
	}

	var t models.TaskApi

	if err := c.BodyParser(&t); err != nil {
		return sendError(
			"Invalid request data",
			fiber.StatusBadRequest,
		)
	}

	if t.ID < 1 {
		return sendError(
			"Task ID is required field",
			fiber.StatusBadRequest,
		)
	}

	user, err := util.GetUserByLocal(c)

	if err != nil {
		return sendError(
			"Cannot find user",
			fiber.StatusBadRequest,
		)
	}

	var task models.Task
	result := db.DB.Where(
		"id = ? AND owner_id = ?", t.ID, user.ID,
	).Model(models.Task{}).First(&task)

	if result.Error != nil {
		return sendError(
			"Cannot find the Task",
			fiber.StatusForbidden,
		)
	}

	task.Title = t.Title
	task.Description = t.Description
	task.Status = t.Status

	result = db.DB.Save(&task)

	if result.Error != nil {
		return sendError(
			"Cannot update task "+result.Error.Error(),
			fiber.StatusForbidden,
		)
	}

	return c.Status(fiber.StatusOK).JSON(task)
}
