package psql 

import (
	model "github.com/andream16/go-storm/model/database"
	"reflect"
	"strings"
	"fmt"
)
var tables = map[string]interface{} {
	"Manufacturer" : model.Manufacturer{},
	"Trend"        : model.Trend{},
	"Item" 		   : model.Item{},
	"Price"        : model.Price{},
	"Review"       : model.Review{},
	"Category"     : model.Category{},
	"Forecast" 	   : model.Forecast{},
	"Currency"     : model.Currency{},
}

type StormTag struct {
	Key         string
	References  string
	Constraint  string
}

type Table struct {
	Name    string
	Fields  []string
	Types   []interface{}
	Options []Option
}

type Option struct {
	Name string
	Value interface{}
}

func CreateTables() {
	var finalTables []Table
	for tableName := range tables {
		var fields    []string
		var types     []interface{}
		var stormTags []Option
		currentModel := tables[tableName]
		typeOfModel := reflect.TypeOf(currentModel)
		for i := 0; i < typeOfModel.NumField(); i++ {
			finalTables[i].Name   = strings.ToLower(tableName)
			finalTables[i].Fields = append(fields, strings.ToLower(typeOfModel.Field(i).Name))
			finalTables[i].Types  = append(types, typeOfModel.Field(i).Type)
			tagsEntries := strings.Split(typeOfModel.Field(i).Tag.Get("storm"), ";")
			for t := range tagsEntries {
				var o Option
				s := strings.Split(tagsEntries[t], ":")
				o.Name = s[0]
				o.Value = s[1]
				stormTags = append(stormTags, o)
			}
			finalTables[i].Options = stormTags
		}
	}
	for k := range finalTables {
		fmt.Println(finalTables[k].Name, finalTables[k].Fields, finalTables[k].Types, finalTables[k].Options)
	}
}