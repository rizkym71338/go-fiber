package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizkym71338/go-fiber/handlers"
	"github.com/rizkym71338/go-fiber/libs"
	"github.com/rizkym71338/go-fiber/modules/user/repository"
)

func Create(c *fiber.Ctx) error {
	err, invalid := repository.Create(c)

	if invalid {
		return handlers.BadRequest(c)
	}

	if libs.IsError(err) {
		return handlers.InternalServerError(c)
	}

	return handlers.Success(c)
}
