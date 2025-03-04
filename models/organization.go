package models

import (
	"time"

	"gorm.io/gorm"
)

type Organization struct {
	ID        uint           `json:"id,omitempty,omitzero" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at,omitempty,omitzero"`
	UpdatedAt time.Time      `json:"updated_at,omitempty,omitzero"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty,omitzero" gorm:"index"`
	Name      string         `json:"name,omitempty,omitzero" gorm:"not null" validate:"required,min=3,max=50"`
	Email     string         `json:"email,omitempty,omitzero" gorm:"unique;not null" validate:"required,email"`
	Phone     string         `json:"phone,omitempty,omitzero" gorm:"unique;not null" validate:"required,min=8,max=8"`
	Register  string         `json:"register,omitempty,omitzero" gorm:"unique;not null" validate:"required"`
	Image     *string        `json:"image,omitempty,omitzero"`
	Deleted   bool           `json:"deleted,omitempty,omitzero"`
}
