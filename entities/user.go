package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string    `gorm:"not null" json:"name" form:"name"`
	Email       string    `gorm:"unique;not null" json:"email" form:"email"`
	Password    string    `gorm:"not null" json:"password" form:"password"`
	PhoneNumber string    `gorm:"not null" json:"phone_number" form:"phone_number"`
	Product     []Product `gorm:"foreignKey:UserID;references:ID"`
	Cart        []Cart    `gorm:"foreignKey:UserID;references:ID"`
	Address     []Address `gorm:"foreignKey:UserID;references:ID"`
	Order       []Order   `gorm:"foreignKey:UserID;references:ID"`
}
