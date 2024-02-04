package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Success(c *fiber.Ctx, m ...string) error {
	message := "Success"
	if len(m) > 0 && m[0] != "" {
		message = m[0]
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": message,
	})
}
