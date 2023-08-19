package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID           uint `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time
	Title        string `json:"title" gorm:"not null"`
	Description  string `json:"description" gorm:"not null"`
	Amount       int    `json:"amount" gorm:"not null"`
	SerialNumber string `json:"serial_number" gorm:"not null"`
	OrderID      uint
}
