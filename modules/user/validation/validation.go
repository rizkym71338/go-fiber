package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/rizkym71338/go-fiber/modules/user/model"
)

var validate *validator.Validate

type Schema struct {
	Name  *string `json:"name" validate:"required"`
	Email *string `json:"email" validate:"required,email"`
}

func Validate(user *model.User) error {
	validate = validator.New()
	schema := &Schema{
		Name:  user.Name,
		Email: user.Email,
	}
	err := validate.Struct(schema)
	return err
}
