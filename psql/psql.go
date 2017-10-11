package psql 

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	model "github.com/andream16/go-storm/model/database"
)

var psql = gorm.DB{}

var tables = map[string]interface{} {
	"Item" 		   : &model.Item{},
	"Price"        : &model.Price{},
	"Review"       : &model.Review{},
	"Category"     : &model.Category{},
	"Forecast" 	   : &model.Forecast{},
	"Manufacturer" : &model.Manufacturer{},
	"Trend"        : &model.Trend{},
	"Currency"     : &model.Currency{},
}

func InitializePsql() (*gorm.DB, error) {
	db, dbErr := gorm.Open("postgres", "host=localhost user=postgres dbname=priceprobe sslmode=disable"); if dbErr != nil {
		return db, dbErr
	}
	psql = *db
	CreateTables()
	defer db.Close()
	return db, nil
}

func GetDbInstance() *gorm.DB {
	return &psql
}

func CreateTables() {
	for k, t := range tables {
		if !psql.HasTable(k) {
			psql.CreateTable(t)
		}
	}
	psql.Model(&model.Price{}).Related(&model.Item{})
}