package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	LastName       string       `json:"last_name,omitempty" gorm:"not null" validate:"required,min=3,max=50"`
	FirstName      string       `json:"first_name,omitempty" gorm:"not null" validate:"required,min=3,max=50"`
	Image          *string      `json:"image,omitempty"`
	Phone          string       `json:"phone,omitempty" gorm:"not null" validate:"required,min=8,max=8"`
	Email          string       `json:"email,omitempty" gorm:"unique;not null" validate:"required,email"`
	Password       string       `json:"password,omitempty" gorm:"not null" validate:"required,min=8"`
	RoleID         uint         `json:"role_id,omitempty" gorm:"not null" validate:"required"`
	Role           Role         `json:"role,omitempty" gorm:"foreignKey:RoleID" validate:"-"`
	OrganizationID uint         `json:"organization_id,omitempty" gorm:"not null" validate:"required"`
	Organization   Organization `json:"organization,omitempty" gorm:"foreignKey:OrganizationID" validate:"-"`
}
