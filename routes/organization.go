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
	protect.Post("", handler.CreateOrganizationHandler)
	protect.Get("", handler.GetOrganizationHandler)
	protect.Patch("/:id", handler.UpdateOrganizationHandler)
	protect.Delete("/:id", handler.DeleteOrganizationHandler)
	protect.Get("/:id", handler.GetOrganizationByIDHandler)
}
