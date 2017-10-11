package database

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Item []Item  `gorm:"ForeignKey:Item"`
	Name string
}
