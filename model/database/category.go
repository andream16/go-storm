package database

type Category struct {
	ID uint `gorm:"primary_key"`
	Item []Item `gorm:"ForeignKey:Item;AssociationForeignKey:Refer"`
	Name string
}
