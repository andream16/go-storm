package psql

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/go-errors/errors"
	"github.com/andream16/go-storm/configuration"
	"github.com/andream16/go-storm/psql/sequence"
	"github.com/andream16/go-storm/psql/table"
)

// Initializes connection to Postgresql client and creates tables.
func InitializePostgresql(conf *configuration.Configuration) error {
	fmt.Println(fmt.Sprintf("Connecting to Postgresql with credentials: user=%s dbname=%s sslmode=%s ...", conf.Database.USER, conf.Database.DBNAME, conf.Database.SSLMODE))
	db, dbErr := sql.Open(conf.Database.DRIVERNAME, fmt.Sprintf("user=%s dbname=%s sslmode=%s",
				        conf.Database.USER, conf.Database.DBNAME, conf.Database.SSLMODE)); if dbErr != nil {
		fmt.Println("Unable to connect to Postgresql, got error: ", errors.New(dbErr))
		return dbErr
	}
	fmt.Println("Successfully connected to Postgresql. Now creating sequences ...")
	createSequences(db)
	fmt.Println("Sequences done. Now creating tables . . .")
	createTables(db)
	fmt.Println("Tables done. Postgresql initialization done.")
	return nil
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
	tableNames := table.TABLES
	tableQueries := table.CREATETABLES
	for _, v := range order {
		_, tableError := db.Exec(tableQueries[tableNames[v]]); if tableError != nil {
			fmt.Println(fmt.Sprintf("unable to create table %s, error: %s", v, tableError))
		} else {
			fmt.Println(fmt.Sprintf("created table %s", v))
		}
	}
}


