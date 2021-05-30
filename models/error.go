package models

import "github.com/gofiber/fiber/v2"

type ApiError struct {
	Error   bool   `json:"error" default:"true"`
	Message string `json:"message"`
}

func DefaultError(msg string) ApiError {
	return ApiError{
		Error:   true,
		Message: msg,
	}
}

func (e ApiError) SendStatus(c *fiber.Ctx, s int) error {
	return c.Status(s).JSON(e)
}
