package libs

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rizkym71338/go-fiber/handlers"
)

func Pagination(c *fiber.Ctx) (int, int) {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		handlers.BadRequest(c, "Invalid page number")
		return 0, 0
	}

	perPage, err := strconv.Atoi(c.Query("perPage", "10"))
	if err != nil || perPage < 1 {
		handlers.BadRequest(c, "Invalid perPage value")
		return 0, 0
	}

	limit := perPage
	offset := (page - 1) * perPage

	return limit, offset
}
