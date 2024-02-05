package repository

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizkym71338/go-fiber/database"
	"github.com/rizkym71338/go-fiber/modules/user/model"
)

func Update(ctx *fiber.Ctx) (error, int64, bool) {
	id := ctx.Params("id")
	var user model.User
	if err := ctx.BodyParser(&user); err != nil {
		return err, 0, true
	}
	result := database.Model.Where("id = ?", id).Updates(&user)
	return result.Error, result.RowsAffected, false
}
