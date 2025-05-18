package models

import "gorm.io/gorm"

type ProductType struct {
	gorm.Model
	Name     string       `json:"name,omitempty,omitzero" gorm:"not null" validate:"required,min=3,max=50"`
	ParentID *uint        `json:"parent_id,omitempty,omitzero" validate:"required"`
	Parent   *ProductType `json:"parent,omitempty,omitzero" gorm:"foreignKey:ParentID"`
}
