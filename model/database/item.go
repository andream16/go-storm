package database

type Item struct {
	Item         string `gorm:"primary_key"`
	Manufacturer string `gorm:"ForeignKey:Manufacturer;AssociationForeignKey:Refer"`
	Category     []Category `gorm:"ForeignKey:Category;AssociationForeignKey:Refer"`
	Reviews      []Review `gorm:"ForeignKey:Review;AssociationForeignKey:Refer"`
	Prices       []Price `gorm:"ForeignKey:Price;AssociationForeignKey:Refer"`
	Forecast     []Forecast `gorm:"ForeignKey:Forecast;AssociationForeignKey:Refer"`
	URL 		 string
	Image 		 string
	Title 	     string
	Description  string
}
