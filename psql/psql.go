package psql

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/go-errors/errors"
	"github.com/andream16/go-storm/configuration"
	"github.com/andream16/go-storm/psql/sequence"
	"github.com/andream16/go-storm/psql/table"
	"github.com/andream16/go-storm/psql/insert"
	"github.com/andream16/go-storm/psql/enum"
)

// Initializes connection to Postgresql client and creates tables.
func InitializePostgresql(conf *configuration.Configuration) (*sql.DB, error) {
	host := conf.Server.Host
	/*if (conf.Environment == "Production") {
		host = "db"
	}*/
	connString := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s", host, conf.Database.USER, conf.Database.DBNAME, conf.Database.SSLMODE)
	fmt.Println(fmt.Sprintf("Connecting to Postgresql with credentials: user=%s dbname=%s sslmode=%s using connection string = %s", conf.Database.USER, conf.Database.DBNAME, conf.Database.SSLMODE, connString))
	db, dbErr := sql.Open(conf.Database.DRIVERNAME, connString); if dbErr != nil {
		fmt.Println("Unable to connect to Postgresql, got error: ", errors.New(dbErr))
		return db, dbErr
	}
	fmt.Println("Successfully connected to Postgresql. Now creating sequences ...")
	createSequences(db)
	fmt.Println("Enum done. Now creating enum . . .")
	createEnum(db)
	fmt.Println("Enum done. Now creating tables . . .")
	createTables(db)
	fmt.Println("Tables done.")
	insertType := insert.DEV_INSERTS
	if (conf.Environment == "Production") {
		insertType = insert.PROD_INSERTS
	}
	fmt.Println("Now inserting default inserts . . .")
	defaultInserts(db, insertType)
	fmt.Println("Inserts done. Postgresql initialization done.")
	return db, nil
}

func createSequences(db *sql.DB) {
	sequences := sequence.SEQUENCES
	for k := range sequences {
		_, sequenceError := db.Query(sequences[k]); if sequenceError != nil {
			fmt.Println(fmt.Sprintf("unable to create sequence for %s, error: %s", k, sequenceError))
		} else {
			fmt.Println(fmt.Sprintf("created sequence for %s", k))
		}
	}
}

func createTables(db *sql.DB) {
	order := table.ORDER
	tableQueries := table.CREATETABLES
	for _, v := range order {
		_, tableError := db.Exec(tableQueries[v]); if tableError != nil {
			fmt.Println(fmt.Sprintf("unable to create table %s, error: %s", v, tableError))
		} else {
			fmt.Println(fmt.Sprintf("created table %s", v))
		}
	}
}

func defaultInserts(db *sql.DB, insertQueries []string) {
	for k := range insertQueries {
		_, insertError := db.Exec(insertQueries[k]); if insertError != nil {
			fmt.Println(fmt.Sprintf("unable to insert row %s, error: %s", insertQueries[k], insertError))
		} else {
			fmt.Println(fmt.Sprintf("inserted row %s", insertQueries[k]))
		}
	}
}

func createEnum(db *sql.DB){
	enums:= enum.ENUMERATIONS
	for k := range enums {
		_, insertError := db.Exec(enums[k]); if insertError != nil {
			fmt.Println(fmt.Sprintf("unable to insert row %s, error: %s", enums[k], insertError))
		} else {
			fmt.Println(fmt.Sprintf("inserted row %s", enums[k]))
		}
	}
}


