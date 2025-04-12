package models

import "time"

type LookUpType struct {
	ID        uint      `json:"id,omitempty,omitzero" gorm:"primarykey"`
	Name      string    `json:"name,omitempty,omitzero" gorm:"not null;unique" validate:"required"`
	CreatedAt time.Time `json:"created_at,omitempty,omitzero"`
	UpdatedAt time.Time `json:"updated_at,omitempty,omitzero"`
	Deleted   bool      `json:"deleted,omitempty,omitzero" gorm:"not null;default:false"`
	DeletedAt time.Time `json:"deleted_at,omitempty,omitzero" gorm:"null"`
}

type LookUpValue struct {
	Code         string     `json:"code,omitempty,omitzero" gorm:"primarykey" validate:"required"`
	LookUpType   LookUpType `json:"lookup_type,omitempty,omitzero" gorm:"foreignKey:LookUpTypeID" validate:"-"`
	LookUpTypeID uint       `json:"lookup_type_id,omitempty,omitzero" gorm:"not null" validate:"required"`
	Meaning      string     `json:"meaning,omitempty,omitzero" gorm:"not null" validate:"required"`
	CreatedAt    time.Time  `json:"created_at,omitempty,omitzero"`
	UpdatedAt    time.Time  `json:"updated_at,omitempty,omitzero"`
	Deleted      bool       `json:"deleted,omitempty,omitzero" gorm:"not null;default:false"`
	DeletedAt    time.Time  `json:"deleted_at,omitempty,omitzero" gorm:"null"`
}
