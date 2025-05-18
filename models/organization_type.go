package models

type OrganizationType struct {
	ID   uint   `json:"id,omitempty,omitzero" gorm:"primaryKey"`
	Name string `json:"name,omitempty,omitzero" gorm:"unique;not null" validate:"required,min=3,max=50"`
}
