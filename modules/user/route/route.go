package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizkym71338/go-fiber/modules/user/controller"
)

func Route(r fiber.Router) {
	route := r.Group("/users")

	route.Get("/", controller.FindMany)
	route.Get("/:id", controller.FindFirst)
	route.Post("/", controller.Create)
	route.Put("/:id", controller.Update)
	route.Delete("/:id", controller.Delete)
}
