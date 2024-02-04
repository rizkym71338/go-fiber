package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizkym71338/go-fiber/handlers"
	"github.com/rizkym71338/go-fiber/libs"
	"github.com/rizkym71338/go-fiber/modules/user/repository"
)

func Delete(c *fiber.Ctx) error {
	err, rowsAffected := repository.Delete(c)

	if libs.IsNoRowsAffected(rowsAffected) {
		return handlers.NotFound(c)
	}

	if libs.IsError(err) {
		return handlers.InternalServerError(c)
	}

	return handlers.Success(c)
}
