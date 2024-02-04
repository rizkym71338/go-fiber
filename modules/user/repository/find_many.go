package repository

import (
	"github.com/rizkym71338/go-fiber/database"
	"github.com/rizkym71338/go-fiber/modules/user/model"
)

func FindMany() ([]model.User, error) {
	var users []model.User
	result := database.Model.Find(&users)
	return users, result.Error
}
