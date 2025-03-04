package models

import "time"

type Product struct {
	CreatedAt     time.Time `json:"created_at,omitzero"`
	UpdatedAt     time.Time `json:"updated_at,omitzero"`
	ID            uint      `json:"id,omitempty" gorm:"primarykey"`
	Name          string
	ProductType   ProductType
	ProductTypeID uint
}
