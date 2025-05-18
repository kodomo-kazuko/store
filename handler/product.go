package handler

import (
	"store/helper"
	"store/query"
	"store/shared"

	"github.com/gofiber/fiber/v2"
)

const resourceName = "product"

func GetProductHandler(c *fiber.Ctx) error {
	_product := query.Product

	products, err := _product.
		WithContext(c.Context()).
		Scopes(helper.Paginate(c)).
		Find()
	if err != nil {
		shared.NotFound(c, resourceName)
	}

	return shared.Found(c, resourceName, products, nil)
}
