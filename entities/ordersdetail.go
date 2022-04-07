package entities

import "gorm.io/gorm"

type OrdersDetail struct {
	gorm.Model
	CartID  *uint `json:"cart_id" form:"cart_id"`
	OrderID *uint `json:"order_id" form:"order_id"`
	Cart    Cart  `gorm:"foreignKey:CartID;references:ID"`
	Order   Order `gorm:"foreignKey:OrderID;references:ID"`
}
