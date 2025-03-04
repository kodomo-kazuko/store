package shared

import (
	"store/form"

	"github.com/gofiber/fiber/v2"
)

func Success(c *fiber.Ctx, status int, message string, data interface{}, rows interface{}) error {
	c.Next()
	c.Status(status)
	return c.JSON(form.Response{
		Success:  true,
		Message:  message,
		Items:    data,
		MetaData: form.MetaData{Total: rows},
	})
}

func Error(c *fiber.Ctx, status int, message string, err error) error {
	c.Status(status)
	return c.JSON(form.Response{
		Success: false,
		Message: message,
		Error:   err,
	})
}

func BadRequest(c *fiber.Ctx, message string) error {
	c.Status(fiber.StatusBadRequest)
	return c.JSON(form.Response{
		Success: false,
		Message: message,
	})
}

func Created(c *fiber.Ctx, resource string, item interface{}) error {
	message := resource + " created successfully"
	return Success(c, fiber.StatusCreated, message, item, nil)
}

func NotCreated(c *fiber.Ctx, resource string, err error) error {
	message := resource + " not created"
	return Error(c, fiber.StatusNotFound, message, err)
}

func Found(c *fiber.Ctx, resource string, item interface{}, row interface{}) error {
	message := resource + " found successfully"
	return Success(c, fiber.StatusOK, message, item, row)
}

func Updated(c *fiber.Ctx, resource string, item interface{}) error {
	message := resource + " updated successfully"
	return Success(c, fiber.StatusOK, message, item, nil)
}

func Deleted(c *fiber.Ctx, resource string) error {
	message := resource + " deleted successfully"
	return Success(c, fiber.StatusOK, message, nil, nil)
}

func NotFound(c *fiber.Ctx, resource string) error {
	message := resource + " not found"
	return Error(c, fiber.StatusNotFound, message, nil)
}

func InternalServerError(c *fiber.Ctx, err error) error {
	return Error(c, fiber.StatusInternalServerError, "an error occured", err)
}

func Signin(c *fiber.Ctx, Token string) error {
	c.Status(fiber.StatusOK)
	return c.JSON(form.Response{Success: true, Message: "Signed in successfully", Token: Token})
}
