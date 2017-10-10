package database

import "time"

type Forecast struct {
	ID uint `gorm:"primary_key";gorm:"AUTO_INCREMENT"`
	Item Item `gorm:"ForeignKey:Item;AssociationForeignKey:Refer"`
	Price float64
	Date time.Time
}
