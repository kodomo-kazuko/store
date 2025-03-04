package handler

import (
	"store/form"
	"store/helper"
	"store/middleware"
	"store/query"
	"store/shared"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

var UserResource = "user"

func LoginUserHandler(c *fiber.Ctx) error {
	var body form.LoginForm

	err := helper.Validation(c, &body)
	if err != nil {
		return shared.BadRequest(c, err.Error())
	}

	query := query.User
	user, err := query.WithContext(c.Context()).
		Where(query.Email.Eq(body.Identifier)).
		Or(query.Phone.Eq(body.Identifier)).
		First()
	if err != nil {
		return shared.NotFound(c, UserResource)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return shared.BadRequest(c, "Invalid credentials")
	}

	// Generate JWT token
	claims := map[string]interface{}{
		"user_id":         user.ID,
		"email":           user.Email,
		"name":            user.FirstName,
		"organization_id": user.OrganizationID,
		"role_id":         user.RoleID,
		"position_id":     user.PositionID,
	}
	token, err := middleware.GenerateJWT(claims)
	if err != nil {
		return shared.InternalServerError(c, err)
	}

	return shared.Signin(c, token)
}
