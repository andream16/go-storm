package database

import "github.com/jinzhu/gorm"

type Price struct {
	gorm.Model
	Item  Item     `gorm:"ForeignKey:Item"`
	Price float64
	Date  string
}