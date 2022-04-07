package entities

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID       uint           `gorm:"not null" json:"user_id" form:"user_id"`
	ProductID    uint           `gorm:"not null" json:"product_id" form:"product_id"`
	Qty          uint           `gorm:"not null" json:"qty" form:"qty"`
	SubTotal     uint           `gorm:"not null" json:"sub_total" form:"sub_total"`
	Product      Product        `gorm:"foreignKey:ProductID;references:ID" json:"product" form:"product"`
	OrdersDetail []OrdersDetail `gorm:"foreignKey:CartID;references:ID"`
}
