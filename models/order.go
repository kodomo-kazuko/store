package models

import "time"

type Order struct {
	ID        uint      `json:"id,omitempty" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at,omitzero"`
	UpdatedAt time.Time `json:"updated_at,omitzero"`
	Product   Product   `json:"product,omitempty" gorm:"foreignKey:ProductID" validate:"-"`
	ProductID uint      `json:"product_id,omitempty" gorm:"not null" validate:"required"`
	User      User      `json:"user,omitempty" gorm:"foreignKey:UserID" validate:"-"`
	UserID    uint      `json:"user_id,omitempty" gorm:"not null" validate:"required"`
}
