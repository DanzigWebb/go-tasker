package util

import (
	"github.com/gofiber/fiber/v2"
	"task-app/db"
	"task-app/models"
)

func GetUserByLocal(c *fiber.Ctx) (*models.User, error) {
	id := c.Locals("id")
	u := new(models.User)
	if res := db.DB.Where("id = ?", id).First(&u); res.RowsAffected <= 0 {
		return nil, res.Error
	}

	return u, nil
}
