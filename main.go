package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rizkym71338/go-fiber/database"
	"github.com/rizkym71338/go-fiber/handlers"
	"github.com/rizkym71338/go-fiber/modules"
)

func main() {
	database.Connect()

	app := fiber.New()
	app.Use(recover.New())

	modules.Routes(app.Group("/api/v1"))

	app.Get("/", func(ctx *fiber.Ctx) error {
		return handlers.Success(ctx)
	})

	app.Listen(":3000")
}
