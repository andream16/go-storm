package database

type Category struct {
	Name string  `storm:"key:primary_key"`
	Item string  `storm:"key:foreign_key;references:Item.Item"`
}
