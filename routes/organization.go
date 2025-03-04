package routes

import (
	"fmt"
	"store/handler"
	"store/middleware"

	"github.com/gofiber/fiber/v2"
)

func organization(router fiber.Router) {
	route := "organization"
	organization := router.Group(fmt.Sprintf("/%s", route))
	protect := organization.Use(middleware.JWTMiddleware())
	protect.Get("", handler.CreateOrganizationHandler)
}
