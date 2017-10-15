package main

import (
	"fmt"
	"os"
	"github.com/andream16/go-storm/psql"
	"github.com/andream16/go-storm/configuration"
	"github.com/andream16/go-storm/endpoint"
)

func main() {
	fmt.Println("Starting Go storm . . .")
	fmt.Println("Setting up configuration . . .")
	conf := configuration.InitConfiguration()
	fmt.Println("Successfully got configuration! Setting up Postgresql . . .")
	db, pgErr := psql.InitializePostgresql(&conf); if pgErr != nil {
		os.Exit(1)
	}
	endpoint.InitializeEndpoint(&conf, db)
}
