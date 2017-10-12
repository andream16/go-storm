package database

type Trend struct {
	ID           uint 	    `storm:"key:primary_key"`
	Manufacturer string		`storm:"key:foreign_key;references:Manufacturer.Name"`
	Date 		 string
	Value		 float64
}
