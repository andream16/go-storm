package database

type Item struct {
	Item         	 string		`storm:"key:primary_key"`
	Manufacturer	 string 	`storm:"key:foreign_key;references:Manufacturer.Name"`
	URL 		 	 string
	Image 		 	 string
	Title 	     	 string
	Description  	 string
}
