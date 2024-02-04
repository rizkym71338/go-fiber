package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func NotFound(c *fiber.Ctx, m ...string) error {
	message := "Not Found"
	if len(m) > 0 && m[0] != "" {
		message = m[0]
	}

	return c.Status(http.StatusNotFound).JSON(fiber.Map{
		"message": message,
	})
}
