package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func InternalServerError(c *fiber.Ctx, m ...string) error {
	message := "Internal Server Error"
	if len(m) > 0 && m[0] != "" {
		message = m[0]
	}

	return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
		"message": message,
	})
}
