package models

import "time"

type Product struct {
	ID           uint `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time
	Title        string `json:"title" gorm:"not null"`
	Description  string `json:"description" gorm:"not null"`
	Amount       int    `json:"amount" gorm:"not null"`
	SerialNumber string `json:"serial_number" gorm:"not null"`
	OrderID      uint
}
