package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func Success(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "Success",
	})
}
