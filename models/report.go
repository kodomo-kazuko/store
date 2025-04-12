package models

import (
	"gorm.io/gorm"
)

type Report struct {
	gorm.Model
	User      User        `json:"user,omitempty" gorm:"foreignKey:UserID" validate:"-"`
	UserID    uint        `json:"user_id,omitempty" gorm:"not null" validate:"required"`
	Product   Product     `json:"product,omitempty" gorm:"foreignKey:ProductID" validate:"-"`
	ProductID uint        `json:"product_id,omitempty" gorm:"not null" validate:"required"`
	Reason    string      `json:"reason,omitempty" gorm:"not null" validate:"required,min=3,max=255"`
	Status    LookUpValue `json:"status,omitempty,omitzero" gorm:"foreignKey:StatusID" validate:"-"`
	StatusID  string      `json:"status_id,omitempty" gorm:"not null" validate:"required"`
}
