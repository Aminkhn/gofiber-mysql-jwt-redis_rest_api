package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	Family    string `json:"family"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	CreatedAt time.Time
	Orders    []Order
}
