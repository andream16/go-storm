package database

import "github.com/jinzhu/gorm"

type Trend struct {
	gorm.Model
	Manufacturer Manufacturer
	Date 		 string
	Value		 float64
}
