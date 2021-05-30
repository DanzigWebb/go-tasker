package router

import (
	"github.com/gofiber/fiber/v2"
	"task-app/db"
	"task-app/models"
	"task-app/util"
)

var sendError = func(c *fiber.Ctx, m string, s int) error {
	return models.DefaultError(m).SendStatus(c, s)
}

func setupTasksRoutes() {
	TASKS.Use(util.SecureAuth())
	TASKS.Get("/", handleGetTasks)
	TASKS.Post("/", handleCreateTask)
	TASKS.Patch("/", handleUpdateTask)
}

func handleGetTasks(c *fiber.Ctx) error {
	c.Accepts("application/json")
	c.Accepts("json", "text")

	u, err := util.GetUserByLocal(c)

	if err != nil {
		return sendError(
			c,
			"Cannot find user by token",
			fiber.StatusForbidden,
		)
	}

	var tasks []models.Task
	result := db.DB.Where("user_id = ?", u.ID).Model(models.Task{}).Find(&tasks)

	if result.Error != nil {
		return sendError(
			c,
			"Cannot find user's tasks",
			fiber.StatusForbidden,
		)
	}

	var response []models.TaskApi

	for _, t := range tasks {
		task := models.TaskApi{
			ID:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			Status:      t.Status,
			CreatedAt:   t.CreatedAt.String(),
			UpdatedAt:   t.UpdatedAt.String(),
		}
		response = append(response, task)
	}

	return c.Status(fiber.StatusOK).JSON(response)

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
		UserID:      u.ID,
	}

	result := db.DB.Create(&task).Model(models.Task{})

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiError{
			Message: result.Error.Error(),
		})
	}

	t.ID = task.ID
	t.CreatedAt = task.CreatedAt.String()
	t.UpdatedAt = task.UpdatedAt.String()

	return c.Status(fiber.StatusOK).JSON(t)
}

func handleUpdateTask(c *fiber.Ctx) error {
	c.Accepts("application/json")
	c.Accepts("json", "text")

	var t models.TaskApi

	if err := c.BodyParser(&t); err != nil {
		return sendError(
			c,
			"Invalid request data",
			fiber.StatusBadRequest,
		)
	}

	if t.ID < 1 {
		return sendError(
			c,
			"Task ID is required field",
			fiber.StatusBadRequest,
		)
	}

	user, err := util.GetUserByLocal(c)

	if err != nil {
		return sendError(
			c,
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
			c,
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
			c,
			"Cannot update task "+result.Error.Error(),
			fiber.StatusForbidden,
		)
	}

	return c.Status(fiber.StatusOK).JSON(task)
}
