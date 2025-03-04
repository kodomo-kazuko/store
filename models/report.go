package models

import (
	"time"

	"gorm.io/gorm"
)

type Report struct {
	ID        uint           `json:"id,omitempty,omitzero" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at,omitempty,omitzero"`
	UpdatedAt time.Time      `json:"updated_at,omitempty,omitzero"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty,omitzero" gorm:"index"`
	User      User           `json:"user,omitempty" gorm:"foreignKey:UserID" validate:"-"`
	UserID    uint           `json:"user_id,omitempty" gorm:"not null" validate:"required"`
	Product   Product        `json:"product,omitempty" gorm:"foreignKey:ProductID" validate:"-"`
	ProductID uint           `json:"product_id,omitempty" gorm:"not null" validate:"required"`
	Reason    string         `json:"reason,omitempty" gorm:"not null" validate:"required,min=3,max=255"`
	Status    LU_LookUpValue `json:"status,omitempty,omitzero" gorm:"foreignKey:Code" validate:"-"`
}
