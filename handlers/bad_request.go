package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func BadRequest(c *fiber.Ctx, m ...string) error {
	message := "Bad Request"
	if len(m) > 0 && m[0] != "" {
		message = m[0]
	}

	return c.Status(http.StatusBadRequest).JSON(fiber.Map{
		"message": message,
	})
}
