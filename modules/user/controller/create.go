package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizkym71338/go-fiber/handlers"
	"github.com/rizkym71338/go-fiber/modules/user/repository"
)

func Create(ctx *fiber.Ctx) error {
	err, invalid := repository.Create(ctx)

	if invalid {
		return handlers.BadRequest(ctx)
	}

	if err != nil {
		return handlers.InternalServerError(ctx)
	}

	return handlers.Success(ctx)
}
