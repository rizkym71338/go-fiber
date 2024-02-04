package modules

import (
	"github.com/gofiber/fiber/v2"
	user "github.com/rizkym71338/go-fiber/modules/user/route"
)

func Routes(r fiber.Router) {
	user.Route(r)
}
