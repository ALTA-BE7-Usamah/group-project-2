package entities

import "gorm.io/gorm"

type CreditCard struct {
	gorm.Model
	Type   string `gorm:"not null" json:"type" form:"type"`
	Name   string `gorm:"not null" json:"name" form:"name"`
	Number string `gorm:"not null" json:"number" form:"number"`
	CVV    uint   `gorm:"not null" json:"cvv" form:"cvv"`
	Month  uint   `gorm:"not null" json:"month" form:"month"`
	Year   uint   `gorm:"not null" json:"year" form:"year"`
}
