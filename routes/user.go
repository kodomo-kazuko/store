package routes

import (
	"fmt"
	"store/handler"

	"github.com/gofiber/fiber/v2"
)

func user(router fiber.Router) {
	route := "user"
	user := router.Group(fmt.Sprintf("/%s", route))

	user.Post("/login", handler.LoginUserHandler)

	// protect := user.Use(middleware.JWTMiddleware())
	// protect.Get("", handler.CreateOrganizationHandler)
}
