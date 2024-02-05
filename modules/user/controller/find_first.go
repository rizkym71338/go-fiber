package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizkym71338/go-fiber/handlers"
	"github.com/rizkym71338/go-fiber/modules/user/repository"
	"gorm.io/gorm"
)

func FindFirst(ctx *fiber.Ctx) error {
	user, err := repository.FindFirst(ctx)

	if err == gorm.ErrRecordNotFound {
		return handlers.NotFound(ctx)
	}

	if err != nil {
		return handlers.InternalServerError(ctx)
	}

	return ctx.JSON(user)
}
