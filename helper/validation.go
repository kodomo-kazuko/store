package helper

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func Validation(c *fiber.Ctx, target interface{}) error {
	if err := c.BodyParser(target); err != nil {

		return err
	}

	validate := validator.New()
	if err := validate.Struct(target); err != nil {
		return err
	}

	return nil
}
