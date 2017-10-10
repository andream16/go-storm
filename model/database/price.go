package database

import "time"


type Price struct {
	ID uint `gorm:"primary_key";gorm:"AUTO_INCREMENT"`
	Price float64
	Date time.Time
	Item Item `gorm:"ForeignKey:Item;AssociationForeignKey:Refer"`
}