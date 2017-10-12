package database

type Currency struct {
	Name    string  `storm:"key:primary_key;"`
	Date    string
	Value   float64
}
