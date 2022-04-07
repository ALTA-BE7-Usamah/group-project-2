package entities

import "gorm.io/gorm"

type OrdersDetail struct {
	gorm.Model
	UserID     uint    `json:"user_id" form:"user_id"`
	OrderID    uint    `json:"order_id" form:"order_id"`
	ProductID  uint    `json:"product_id" form:"product_id"`
	TotalPrice uint    `json:"total_price" form:"total_price"`
	User       User    `gorm:"foreignKey:UserID;references:ID"`
	Order      Order   `gorm:"foreignKey:OrderID;references:ID"`
	Product    Product `gorm:"foreignKey:ProductID;references:ID"`
}
