package models

import "time"

type History struct {
	ID          int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt   time.Time `json:"created_at"`
	DeliveredAt time.Time `json:"delivered_at"`
	Delivered   bool      `json:"delivered" gorm:"default:false"`
	ProductID   int64     `json:"product_id" gorm:"index;not null"`
	Product     Product   `json:"product" gorm:"foreignKey:ProductID;references:ID"`
	UserID      int64     `json:"user_id" gorm:"index;not null"`
	User        User      `json:"user" gorm:"foreignKey:UserID;references:ID"`
}
