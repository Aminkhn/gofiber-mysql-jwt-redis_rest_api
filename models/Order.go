package models

type Order struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	UserId    int     `json:"user_id"`
	user      User    `gorm:"foreignKey:UserId"`
	ProductId int     `json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductId"`
}
