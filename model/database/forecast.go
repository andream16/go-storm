package database

import "time"

type Forecast struct {
	ID uint `gorm:"primary_key";gorm:"AUTO_INCREMENT"`
	Price float64
	Date time.Time
	Item Item `gorm:"ForeignKey:Item;AssociationForeignKey:Refer"`
}
