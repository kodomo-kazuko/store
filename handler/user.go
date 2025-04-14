package handler

import (
	"store/form"
	"store/helper"
	"store/middleware"
	"store/query"
	"store/shared"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gen"
)

var UserResource = "user"

func LoginUserHandler(ctx *fiber.Ctx) error {
	var body form.LoginForm

	err := helper.Validation(ctx, &body)
	if err != nil {
		return shared.BadRequest(ctx, err.Error())
	}

	_user := query.User
	user, err := _user.WithContext(ctx.Context()).
		Where(_user.Email.Eq(body.Identifier)).
		Or(_user.Phone.Eq(body.Identifier)).
		First()
	if err != nil {
		return shared.NotFound(ctx, UserResource)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return shared.BadRequest(ctx, "Invalid credentials")
	}

	// Generate JWT token
	claims := map[string]interface{}{
		"user_id":         user.ID,
		"email":           user.Email,
		"name":            user.FirstName,
		"organization_id": user.OrganizationID,
		"role_id":         user.RoleID,
	}
	token, err := middleware.GenerateJWT(claims)
	if err != nil {
		return shared.InternalServerError(ctx, err)
	}

	return shared.Signin(ctx, token)
}

func GetUserAccountHandler(ctx *fiber.Ctx) error {
	user_id := middleware.ExtractUserID(ctx)
	_user := query.User
	user, err := _user.WithContext(ctx.Context()).
		Where(_user.ID.Eq(user_id)).
		First()
	if err != nil {
		return shared.NotFound(ctx, UserResource)
	}
	return shared.Found(ctx, UserResource, user, nil)
}

func GetUserByIDHandler(ctx *fiber.Ctx) error {
	id := middleware.GetIDFromParams(ctx)

	_user := query.User
	user, err := _user.WithContext(ctx.Context()).
		Where(_user.ID.Eq(id)).
		First()
	if err != nil {
		return shared.NotFound(ctx, UserResource)
	}
	return shared.Found(ctx, UserResource, user, nil)
}

func GetUserHandler(ctx *fiber.Ctx) error {
	_user := query.User

	params := helper.QueryType{
		"email": func(value string) gen.Condition { return _user.Email.Lower().Like(value) },
		"name":  func(value string) gen.Condition { return _user.FirstName.Lower().Like(value) },
		"phone": func(value string) gen.Condition { return _user.Phone.Lower().Like(value) },
	}

	conds := helper.BuildConds(ctx, params)

	users, err := _user.WithContext(ctx.Context()).
		Scopes(
			helper.Where(conds...),
			helper.Paginate(ctx)).
		Find()
	if err != nil {
		return shared.NotFound(ctx, UserResource)
	}
	return shared.Found(ctx, UserResource, users, nil)
}
