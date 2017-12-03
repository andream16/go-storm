package service

import (
	"database/sql"
	"github.com/andream16/go-storm/model/request"
	"fmt"
	"errors"
)

func GetCurrencyByName(name string, db *sql.DB) (request.Currency, error) {
	stmt, err := db.Prepare(`SELECT name,date,value FROM currency WHERE name=$1 ORDER BY date DESC`); if err != nil {
		defer stmt.Close()
		return request.Currency{}, err
	}
	defer stmt.Close()
	rows, queryError := stmt.Query(name); if queryError != nil {
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
	stmt, err := db.Prepare(`INSERT INTO currency(name,date,value) VALUES ($1,$2,$3)`); if err != nil {
		defer stmt.Close()
		return err
	}
	defer stmt.Close()
	name := currencies.Name
	for _, currency := range currencies.Trend {
		_, insertError := stmt.Exec(name, currency.Date,currency.Value); if insertError != nil {
			return insertError
		}
	}
	return nil
}

func EditCurrency(currencies request.Currency, db *sql.DB) error {
	name := currencies.Name
	stmtInsert, err := db.Prepare(`INSERT INTO currency(name,date,value) VALUES ($1,$2,$3)`); if err != nil {
		defer stmtInsert.Close()
		return err
	}
	defer stmtInsert.Close()
	deleteError := DeleteCurrency(name, db); if deleteError != nil {
		return deleteError
	}
	for _, currency := range currencies.Trend {
		_, insertError := stmtInsert.Exec(name, currency.Date,currency.Value)
		if insertError != nil {
			return insertError
		}
	}
	return nil
}

func DeleteCurrency(name string, db *sql.DB) error {
	stmtDelete, err := db.Prepare(`DELETE FROM currency WHERE name=$1`); if err != nil {
		defer stmtDelete.Close()
		return err
	}
	defer stmtDelete.Close()
	_, deleteError := stmtDelete.Exec(name); if deleteError != nil {
		return deleteError
	}
	return nil
}
