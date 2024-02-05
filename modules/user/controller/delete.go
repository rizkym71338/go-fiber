package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizkym71338/go-fiber/handlers"
	"github.com/rizkym71338/go-fiber/modules/user/repository"
)

func Delete(ctx *fiber.Ctx) error {
	err, rowsAffected := repository.Delete(ctx)

	if rowsAffected == 0 {
		return handlers.NotFound(ctx)
	}

	if err != nil {
		return handlers.InternalServerError(ctx)
	}

	return handlers.Success(ctx)
}
