package main

import (
	"fmt"
	"github.com/andream16/go-storm/psql"
	"log"
)

func main() {
	fmt.Println("Starting Go storm . . .")
	_, dbError := psql.InitializePsql(); if dbError != nil {
		log.Fatal(dbError.Error())
	}
}
