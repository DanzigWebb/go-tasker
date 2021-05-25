package util

import (
	"github.com/gofiber/fiber/v2"
	"strings"
)

func GetAccessToken(c *fiber.Ctx) string {
	header := c.Get("Authorization")
	slice := strings.Split(header, "Bearer ")
	return slice[len(slice)-1]
}

func GetRefreshToken(c *fiber.Ctx) (string, error) {
	type refreshReq struct {
		RefreshToken string `json:"refreshToken"`
	}

	refresh := new(refreshReq)
	if err := c.BodyParser(&refresh); err != nil {
		return "", err
	}

	return refresh.RefreshToken, nil
}
