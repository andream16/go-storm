package psql 

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ConnectToPsql() {
	db, dbErr := gorm.Open("postgres", "host=localhost user=postgres sslmode=disable"); if dbErr != nil {
		fmt.Println("Unable to open connection with posgres. Got: ", dbErr.Error())
	}	
	fmt.Println("Successfully connected to postgres.")
	defer db.Close()
}
