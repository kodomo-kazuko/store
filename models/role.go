package models

import "time"

type Role struct {
	CreatedAt time.Time `json:"created_at,omitzero"`
	UpdatedAt time.Time `json:"updated_at,omitzero"`
	ID        uint      `json:"id,omitempty" gorm:"primarykey"`
	Name      string    `json:"name,omitempty" gorm:"unique"`
	Deleted   bool      `json:"deleted,omitempty" gorm:"not null;default:false"`
}
