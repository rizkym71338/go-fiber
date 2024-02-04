package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizkym71338/go-fiber/handlers"
	"github.com/rizkym71338/go-fiber/libs"
	"github.com/rizkym71338/go-fiber/modules/user/repository"
)

func FindMany(c *fiber.Ctx) error {
	users, err := repository.FindMany()

	if len(users) == 0 {
		return handlers.NotFound(c)
	}

	if libs.IsError(err) {
		return handlers.InternalServerError(c)
	}

	return c.JSON(users)
}
