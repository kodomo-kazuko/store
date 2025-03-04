package models

import (
	"time"
)

type User struct {
	CreatedAt      time.Time    `json:"created_at,omitzero"`
	UpdatedAt      time.Time    `json:"updated_at,omitzero"`
	ID             uint         `json:"id,omitempty" gorm:"primarykey"`
	LastName       string       `json:"last_name,omitempty" gorm:"not null" validate:"required,min=3,max=50"`
	FirstName      string       `json:"first_name,omitempty" gorm:"not null" validate:"required,min=3,max=50"`
	Image          *string      `json:"image,omitempty"`
	Phone          string       `json:"phone,omitempty" gorm:"not null" validate:"required,min=8,max=8"`
	Email          string       `json:"email,omitempty" gorm:"unique;not null" validate:"required,email"`
	Password       string       `json:"password,omitempty" gorm:"not null" validate:"required,min=8"`
	PositionID     uint         `json:"position_id,omitempty" gorm:"not null" validate:"required"`
	Position       Position     `json:"position,omitempty" gorm:"foreignKey:PositionID" validate:"-"`
	RoleID         uint         `json:"role_id,omitempty" gorm:"not null" validate:"required"`
	Role           Role         `json:"role,omitempty" gorm:"foreignKey:RoleID" validate:"-"`
	OrganizationID uint         `json:"organization_id,omitempty" gorm:"not null" validate:"required"`
	Organization   Organization `json:"organization,omitempty" gorm:"foreignKey:OrganizationID" validate:"-"`
	Deleted        bool         `json:"deleted,omitempty" gorm:"not null;default:false"`
	Fcm            *string      `json:"fcm,omitempty" gorm:"null" validate:"-"`
	Description    string       `json:"description,omitempty" gorm:"not null" validate:"required,min=3,max=255"`
	// Test           string
}
