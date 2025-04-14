package handler

import (
	"store/helper"
	"store/middleware"
	"store/models"
	"store/query"
	"store/shared"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gen"
)

var resource = "organization"

func CreateOrganizationHandler(ctx *fiber.Ctx) error {

	var organization models.Organization

	err := helper.Validation(ctx, &organization)
	if err != nil {
		return shared.BadRequest(ctx, err.Error())
	}

	_organization := query.Organization

	err = _organization.WithContext(ctx.Context()).Create(&organization)
	if err != nil {
		return shared.InternalServerError(ctx, err)
	}
	return shared.Created(ctx, resource, nil)
}

func GetOrganizationHandler(ctx *fiber.Ctx) error {
	_organization := query.Organization

	params := helper.QueryType{
		"email":    func(value string) gen.Condition { return _organization.Email.Lower().Like(value) },
		"name":     func(value string) gen.Condition { return _organization.Name.Lower().Like(value) },
		"register": func(value string) gen.Condition { return _organization.Register.Lower().Like(value) },
		"phone":    func(value string) gen.Condition { return _organization.Phone.Lower().Like(value) },
		"start_date": func(value string) gen.Condition {
			date, _ := helper.ToDate(value, false)
			return _organization.CreatedAt.Gte(date)
		},
		"end_date": func(value string) gen.Condition {
			date, _ := helper.ToDate(value, true)
			return _organization.CreatedAt.Lte(date)
		},
	}

	conds := helper.BuildConds(ctx, params)

	data, err := _organization.WithContext(ctx.Context()).
		Scopes(helper.Where(conds...), helper.Paginate(ctx)).
		Find()
	if err != nil {
		return shared.InternalServerError(ctx, err)
	}

	return shared.Found(ctx, resource, data, nil)
}

func UpdateOrganizationHandler(ctx *fiber.Ctx) error {
	var body models.Organization

	id := middleware.GetIDFromParams(ctx)

	err := helper.Validation(ctx, &body)
	if err != nil {
		return shared.BadRequest(ctx, err.Error())
	}

	_organization := query.Organization

	data, err := _organization.WithContext(ctx.Context()).
		Where(_organization.ID.Eq(id)).
		Updates(&body)
	if err != nil {
		return shared.InternalServerError(ctx, err)
	}

	return shared.Updated(ctx, resource, data)
}

func DeleteOrganizationHandler(ctx *fiber.Ctx) error {
	id := middleware.GetIDFromParams(ctx)

	_organization := query.Organization

	_, err := _organization.WithContext(ctx.Context()).
		Where(_organization.ID.Eq(id)).
		Delete()
	if err != nil {
		return shared.InternalServerError(ctx, err)
	}

	return shared.Deleted(ctx, resource)
}

func GetOrganizationByIDHandler(ctx *fiber.Ctx) error {
	id := middleware.GetIDFromParams(ctx)

	_organization := query.Organization

	data, err := _organization.WithContext(ctx.Context()).
		Where(_organization.ID.Eq(id)).
		First()
	if err != nil {
		return shared.InternalServerError(ctx, err)
	}

	return shared.Found(ctx, resource, data, nil)
}
