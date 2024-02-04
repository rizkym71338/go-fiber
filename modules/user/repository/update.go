package repository

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizkym71338/go-fiber/database"
	"github.com/rizkym71338/go-fiber/libs"
	"github.com/rizkym71338/go-fiber/modules/user/model"
)

func Update(c *fiber.Ctx) (error, int64, bool) {
	id := c.Params("id")
	var user model.User
	if err := c.BodyParser(&user); libs.IsError(err) {
		return err, 0, true
	}
	result := database.Model.Where("id = ?", id).Updates(&user)
	return result.Error, result.RowsAffected, false
}
