package database

import "github.com/jinzhu/gorm"

type Item struct {
	gorm.Model
	Item         string 	  `gorm:"primary_key"`
	Manufacturer Manufacturer `gorm:"ForeignKey:Name"`
	URL 		 string
	Image 		 string
	Title 	     string
	Description  string
}
