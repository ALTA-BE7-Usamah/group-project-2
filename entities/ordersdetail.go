package entities

import "gorm.io/gorm"

type OrdersDetail struct {
	gorm.Model
	UserID     uint    `json:"user_id" form:"user_id"`
	ProductID  uint    `json:"product_id" form:"product_id"`
	TotalPrice uint    `json:"total_price" form:"total_price"`
	Status     string  `gorm:"default:in process" json:"status_order" form:"status_order"`
	Product    Product `gorm:"foreignKey:ProductID;references:ID"`
}
