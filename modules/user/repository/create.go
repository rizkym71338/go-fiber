package repository

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizkym71338/go-fiber/database"
	"github.com/rizkym71338/go-fiber/libs"
	"github.com/rizkym71338/go-fiber/modules/user/model"
)

func Create(c *fiber.Ctx) (error, bool) {
	var user model.User
	if err := c.BodyParser(&user); libs.IsError(err) {
		return err, true
	}
	result := database.Model.Create(&user)
	return result.Error, false
}
