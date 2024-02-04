package database

import user "github.com/rizkym71338/go-fiber/modules/user/model"

func Migration() {
	Model.AutoMigrate(&user.User{})
}
