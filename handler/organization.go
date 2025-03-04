package handler

import (
	"store/helper"
	"store/models"
	"store/query"
	"store/shared"

	"github.com/gofiber/fiber/v2"
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
	_org := query.Organization

	data, err := _org.WithContext(ctx.Context()).Find()
	if err != nil {
		return shared.InternalServerError(ctx, err)
	}

	return shared.Found(ctx, resource, data, nil)

}
