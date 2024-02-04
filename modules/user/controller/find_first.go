package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizkym71338/go-fiber/handlers"
	"github.com/rizkym71338/go-fiber/libs"
	"github.com/rizkym71338/go-fiber/modules/user/repository"
)

func FindFirst(c *fiber.Ctx) error {
	user, err := repository.FindFirst(c)

	if libs.IsNotFound(err) {
		return handlers.NotFound(c)
	}

	if libs.IsError(err) {
		return handlers.InternalServerError(c)
	}

	return c.JSON(user)
}
