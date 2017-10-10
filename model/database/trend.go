package database

import "time"

type Trend struct {
	ID uint `gorm:"primary_key";gorm:"AUTO_INCREMENT"`
	Manufacturer string `gorm:"ForeignKey:Manufacturer;AssociationForeignKey:Refer"`
	Date time.Time
	Value float64
}
