package models

import (
	"gorm.io/gorm"
)

type Organization struct {
	gorm.Model
	Name     string  `json:"name,omitempty,omitzero" gorm:"not null" validate:"required,min=3,max=50"`
	Email    string  `json:"email,omitempty,omitzero" gorm:"unique;not null" validate:"required,email"`
	Phone    string  `json:"phone,omitempty,omitzero" gorm:"unique;not null" validate:"required,min=8,max=8"`
	Register string  `json:"register,omitempty,omitzero" gorm:"unique;not null" validate:"required"`
	Image    *string `json:"image,omitempty,omitzero" gorm:"default:null"`
	Deleted  bool    `json:"deleted,omitempty,omitzero" gorm:"default:false"`
	Address  string  `json:"address,omitempty,omitzero" gorm:"not null" validate:"required,min=3,max=50"`
}
