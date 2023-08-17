package models

import "time"

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	Family    string `json:"family"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	CreatedAt time.Time
}
