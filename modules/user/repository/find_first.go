package repository

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizkym71338/go-fiber/database"
	"github.com/rizkym71338/go-fiber/modules/user/model"
)

func FindFirst(c *fiber.Ctx) (model.User, error) {
	id := c.Params("id")
	var user model.User
	result := database.Model.Where("id = ?", id).First(&user)
	return user, result.Error
}
