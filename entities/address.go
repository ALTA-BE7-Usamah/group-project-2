package entities

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserID  uint   `gorm:"not null" json:"user_id" form:"user_id"`
	Street  string `gorm:"not null" json:"street" form:"street"`
	City    string `gorm:"not null" json:"city" form:"city"`
	State   string `gorm:"not null" json:"state" form:"state"`
	ZipCode uint   `gorm:"not null" json:"zip_code" form:"zip_code"`
}
