package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID            uint           `json:"id,omitempty,omitzero" gorm:"primarykey"`
	CreatedAt     time.Time      `json:"created_at,omitempty,omitzero"`
	UpdatedAt     time.Time      `json:"updated_at,omitempty,omitzero"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at,omitempty,omitzero" gorm:"index"`
	Name          string         `json:"name,omitempty" gorm:"not null" validate:"required,min=3,max=50"`
	ProductType   LU_LookUpValue `json:"product_type,omitempty,omitzero" gorm:"foreignKey:ProductTypeID"`
	ProductTypeID string         `json:"product_type_id,omitempty,omitzero" gorm:"not null" validate:"required"`
}
