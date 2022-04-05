package entities

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	CartID      []uint     `gorm:"not null" json:"cart_id" form:"cart_id"`
	TotalPrice  uint       `gorm:"not null" json:"total_price" form:"total_price"`
	StatusOrder uint       `gorm:"not null" json:"status_order" form:"status_order"`
	Address     Address    `gorm:"not null" json:"address" form:"address"`
	CreditCard  CreditCard `gorm:"not null" json:"credit_card" form:"credit_card"`
	Cart        []Cart     `gorm:"foreignKey:CartID;references:ID"`
}
