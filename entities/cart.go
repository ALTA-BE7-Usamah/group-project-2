package entities

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID     uint    `gorm:"not null" json:"user_id" form:"user_id"`
	ProductID  uint    `gorm:"not null" json:"product_id" form:"product_id"`
	OrderID    uint    `gorm:"not null" json:"order_id" form:"order_id"`
	Quantity   uint    `gorm:"not null" json:"quantity" form:"quantity"`
	TotalPrice uint    `gorm:"not null" json:"total_price" form:"total_price"`
	Product    Product `gorm:"foreignKey:ProductID;references:ID" json:"product" form:"product"`
}
