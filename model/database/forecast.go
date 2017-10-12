package database

type Forecast struct {
	ID      uint       `storm:"key:primary_key"`
	Item    string     `storm:"key:foreign_key;references:Item.Item"`
	Price   float64
	Date    string
}
