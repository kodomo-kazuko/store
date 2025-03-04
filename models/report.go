package models

import "time"

type Report struct {
	ID        uint      `json:"id,omitempty" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at,omitzero"`
	UpdatedAt time.Time `json:"updated_at,omitzero"`
	User      User      `json:"user,omitempty" gorm:"foreignKey:UserID" validate:"-"`
	UserID    uint      `json:"user_id,omitempty" gorm:"not null" validate:"required"`
	Product   Product   `json:"product,omitempty" gorm:"foreignKey:ProductID" validate:"-"`
	ProductID uint      `json:"product_id,omitempty" gorm:"not null" validate:"required"`
	Reason    string    `json:"reason,omitempty" gorm:"not null" validate:"required,min=3,max=255"`
}
