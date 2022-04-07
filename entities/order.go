package entities

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID       uint           `gorm:"not null" json:"user_id" form:"user_id"`
	TotalPrice   uint           `gorm:"not null" json:"total_price" form:"total_price"`
	StatusOrder  uint           `gorm:"not null" json:"status_order" form:"status_order"`
	Address      Address        `gorm:"foreignKey:ID;references:ID" json:"address" form:"address"`
	CreditCard   CreditCard     `gorm:"foreignKey:ID;references:ID" json:"credit_card" form:"credit_card"`
	OrdersDetail []OrdersDetail `gorm:"foreignKey:OrderID;references:ID"`
}
