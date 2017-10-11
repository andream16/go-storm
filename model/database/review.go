package database

import (
	"github.com/jinzhu/gorm"
)

type Review struct {
	gorm.Model
	Item 	  Item    `gorm:"ForeignKey:Item"`
	Text 	  string
	Date 	  string
	Sentiment uint
	Stars     uint
}