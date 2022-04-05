package entities

import "gorm.io/gorm"

type Catagory struct {
	gorm.Model
	CatagoryName string    `gorm:"not null" json:"catagory_name" form:"catagory_name"`
	Product      []Product `gorm:"foreignKey:CatagoryID;references:ID"`
}
