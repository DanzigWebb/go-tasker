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

	id := c.Locals("id")
	u := new(models.User)
	if res := db.DB.Where("uuid = ?", id).First(&u); res.RowsAffected <= 0 {
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

	return c.Status(fiber.StatusOK).JSON(t)
}

func handleUpdateTask(c *fiber.Ctx) error {
	c.Accepts("application/json")
	c.Accepts("json", "text")

	var t models.TaskApi

	if err := c.BodyParser(&t); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			models.DefaultError("Invalid request data"),
		)
	}

	if t.ID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(
			models.DefaultError("Task ID is required field"),
		)
	}

	user, err := util.GetUserByLocal(c)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			models.DefaultError("Cant find user"),
		)
	}

	var task models.Task
	result := db.DB.Where(
		"id = ? AND owner_id = ?", t.ID, user.ID,
	).Model(models.Task{}).First(&task)

	if result.Error != nil {
		return c.Status(fiber.StatusForbidden).JSON(models.DefaultError("Cannot find the Task"))
	}

	task.Title = t.Title
	task.Description = t.Description
	task.Status = t.Status

	result = db.DB.Save(&task)

	if result.Error != nil {
		return c.Status(fiber.StatusForbidden).JSON(models.DefaultError("Cant update task " + result.Error.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(task)
}
