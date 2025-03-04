package routes

import (
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	app.Route("", func(router fiber.Router) {
		organization(app)
		user(app)
	})
}
