package repository

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizkym71338/go-fiber/database"
	"github.com/rizkym71338/go-fiber/modules/user/model"
)

func Delete(c *fiber.Ctx) (error, int64) {
	id := c.Params("id")
	var user model.User
	result := database.Model.Where("id = ?", id).Delete(&user)
	return result.Error, result.RowsAffected
}
