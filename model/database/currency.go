package database

import "github.com/jinzhu/gorm"

type Currency struct {
	gorm.Model
	Name    string  `gorm:"primary_key"`
	Date    string
	Value   float64
}
