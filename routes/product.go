package routes

import (
	"fmt"
	"store/handler"
	"store/middleware"

	"github.com/gofiber/fiber/v2"
)

func product(router fiber.Router) {
	path := "product"
	route := router.Group(fmt.Sprintf("/%s", path))

	protect := route.Use(middleware.JWTMiddleware())
	protect.Get("/", handler.GetProductHandler)
}
