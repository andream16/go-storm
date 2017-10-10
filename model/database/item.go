package database

type Item struct {
	Item         string `gorm:"primary_key"`
	Category     []string
	Manufacturer string `gorm:"ForeignKey:Manufacturer;AssociationForeignKey:Refer"`
	URL 		 string
	Image 		 string
	Title 	     string
	Description  string
}
