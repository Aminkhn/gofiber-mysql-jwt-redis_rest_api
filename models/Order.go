package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ID     uint `json:"id" gorm:"primaryKey"`
	UserID uint `json:"user_id"`
	//user      User    `gorm:"foreignKey:UserId"`
	//ProductID uint     `json:"product_id"`
	//Product   Product `gorm:"foreignKey:ProductId"`
	Products []Product
}
