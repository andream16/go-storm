package database

type Manufacturer struct {
	Name string `gorm:"primary_key";gorm:"AUTO_INCREMENT"`
	Item []Item `gorm:"ForeignKey:Item;AssociationForeignKey:Refer"`
}
