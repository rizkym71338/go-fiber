package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizkym71338/go-fiber/database"
	"github.com/rizkym71338/go-fiber/modules"
)

func main() {
	database.Connect()

	app := fiber.New()
	api := app.Group("/api/v1")

	modules.Routes(api)

	app.Listen(":3000")
}
