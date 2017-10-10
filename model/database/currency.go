package database

import "time"

type Currency struct {
	ID uint `gorm:"primary_key";gorm:"AUTO_INCREMENT"`
	Name string
	Date time.Time
	Value float64
}
