package database

import "github.com/jinzhu/gorm"

type Forecast struct {
	gorm.Model
	Item    Item     `gorm:"ForeignKey:Item"`
	Price   float64
	Date    string
}
