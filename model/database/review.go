package database

import "time"

type Review struct {
	ID uint `gorm:"primary_key";gorm:"AUTO_INCREMENT"`
	Item Item `gorm:"ForeignKey:Item;AssociationForeignKey:Refer"`
	Text string
	Date time.Time
	Sentiment uint
}