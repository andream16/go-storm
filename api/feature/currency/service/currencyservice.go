package service

import (
	"database/sql"
	"github.com/andream16/go-storm/model/request"
	"fmt"
	"errors"
)

func GetCurrencyByName(name string, db *sql.DB) (request.Currency, error) {
	rows, queryError := db.Query(`SELECT name,date,value FROM currency WHERE name=$1 ORDER BY date DESC`, name); if queryError != nil {
		return request.Currency{}, errors.New(fmt.Sprintf("Unable to get date,values for currency %s. Error: %s", name, queryError.Error()))
	}
	defer rows.Close()
	var currencies request.Currency
	for rows.Next() {
		var currency request.CurrencyEntry
		rowError := rows.Scan(&currencies.Name,&currency.Date,&currency.Value); if rowError != nil {
			return request.Currency{}, errors.New(fmt.Sprintf("Unable to unmarshal date,values for currency %s. Error: %s", name, rowError.Error()))
		}
		currencies.Trend = append(currencies.Trend, currency)
	}; iterationError := rows.Err(); if iterationError != nil {
		return request.Currency{}, errors.New(fmt.Sprintf("No date,values found for currency %s. Error: %s", name, iterationError.Error()))
	}; if len(currencies.Trend) == 0 {
		return request.Currency{}, errors.New(fmt.Sprintf("No date,values found for currency %s", name))
	}
	return currencies, nil
}

func AddCurrencies(currencies request.Currency, db *sql.DB) error {
	name := currencies.Name
	for _, currency := range currencies.Trend {
		_, insertError := db.Query(`INSERT INTO currency(name,date,value) VALUES ($1,$2,$3)`, name, currency.Date,currency.Value)
		if insertError != nil {
			return insertError
		}
	}
	return nil
}

func EditCurrency(currencies request.Currency, db *sql.DB) error {
	name := currencies.Name
	_, deleteError := db.Query(`DELETE FROM currency WHERE name=$1`, name); if deleteError != nil {
		return deleteError
	}
	for _, currency := range currencies.Trend {
		_, insertError := db.Query(`INSERT INTO currency(name,date,value) VALUES ($1,$2,$3)`, name, currency.Date,currency.Value)
		if insertError != nil {
			return insertError
		}
	}
	return nil
}

func DeleteCurrency(name string, db *sql.DB) error {
	_, deleteError := db.Query(`DELETE FROM currency WHERE name=$1`, name); if deleteError != nil {
		return deleteError
	}
	return nil
}
