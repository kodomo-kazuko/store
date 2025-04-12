package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name           string       `json:"name,omitempty" gorm:"not null" validate:"required,min=3,max=50"`
	Price          float64      `json:"price,omitempty" gorm:"not null" validate:"required"`
	Stock          int          `json:"stock,omitempty" gorm:"not null" validate:"required"`
	ProductType    ProductType  `json:"product_type,omitempty,omitzero" gorm:"foreignKey:ProductTypeID"`
	ProductTypeID  uint         `json:"product_type_id,omitempty,omitzero" gorm:"not null" validate:"required"`
	Organization   Organization `json:"organization,omitempty,omitzero" validate:"-"`
	OrganizationID uint         `json:"organization_id,omitempty,omitzero" validate:"required"`
}
