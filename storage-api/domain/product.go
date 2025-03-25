package domain

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `gorm:"type:varchar(128)"`
	Description string `gorm:"type:varchar(256)"`
	Balance     int
}
