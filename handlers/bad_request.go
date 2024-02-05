package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func BadRequest(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.ErrBadRequest)
}
