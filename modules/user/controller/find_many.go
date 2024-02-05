package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizkym71338/go-fiber/handlers"
	"github.com/rizkym71338/go-fiber/modules/user/repository"
)

func FindMany(ctx *fiber.Ctx) error {
	users, err := repository.FindMany()

	if len(users) == 0 {
		return handlers.NotFound(ctx)
	}

	if err != nil {
		return handlers.InternalServerError(ctx)
	}

	return ctx.JSON(users)
}
