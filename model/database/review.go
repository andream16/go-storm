package database

type Review struct {
	ID        uint      `storm:"key:primary_key;"`
	Item 	  string    `storm:"key:foreign_key;references:Item.Item"`
	Text 	  string
	Date 	  string
	Sentiment int       `storm:"constraint:value[1,5]"`
	Stars     int		`storm:"constraint:value[0,1]"`
}
