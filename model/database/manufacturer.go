package database

import "github.com/jinzhu/gorm"

type Manufacturer struct {
	gorm.Model
	Name string `gorm:"primary_key"`
}
