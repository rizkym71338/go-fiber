package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func NotFound(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNotFound).JSON(fiber.ErrNotFound)
}
