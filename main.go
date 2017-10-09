package main

import (
	"fmt"
	"github.com/andream16/go-storm/psql"
)

func main() {
	fmt.Println("Starting Go storm . . .")
	psql.ConnectToPsql()
}
