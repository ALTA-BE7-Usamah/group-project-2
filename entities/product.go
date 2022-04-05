package entities

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	UserID       uint     `gorm:"not null" json:"user_id" form:"user_id"`
	CatagoryID   uint     `gorm:"not null" json:"catagory_id" form:"catagory_id"`
	ProductTitle string   `gorm:"not null" json:"product_title" form:"product_title"`
	Price        uint     `gorm:"not null" json:"price" form:"price"`
	Stock        uint     `gorm:"not null" json:"stock" form:"stock"`
	UrlProduct   string   `gorm:"not null" json:"url_product" form:"url_product"`
	Catagory     Catagory `gorm:"foreignKey:CatagoryID;references:ID"`
	Cart         []Cart   `gorm:"foreignKey:ProductID;references:ID"`
}
