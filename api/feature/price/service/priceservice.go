package service

import (
	"github.com/andream16/go-storm/model/request"
	"database/sql"
	"fmt"
	"errors"
)

func GetPrices(itemId string, db *sql.DB) (request.Prices, error) {
	stmt, err := db.Prepare(`SELECT price,date FROM price WHERE item=$1 ORDER BY date ASC`); if err != nil {
		return request.Prices{}, err
	}
	defer stmt.Close()
	rows, queryError := stmt.Query(itemId); if queryError != nil {
		return request.Prices{}, errors.New(fmt.Sprintf("Unable to get prices for item %s. Error: %s", itemId, queryError.Error()))
	}
	defer rows.Close()
	var prices request.Prices
	prices.Item = itemId
	for rows.Next() {
		var price request.PriceEntry
		rowError := rows.Scan(&price.Price, &price.Date); if rowError != nil {
			return request.Prices{}, errors.New(fmt.Sprintf("Unable to unmarshal prices for item %s. Error: %s", itemId, rowError.Error()))
		}
		prices.Prices = append(prices.Prices, price)
	}; iterationError := rows.Err(); if iterationError != nil {
		return request.Prices{}, errors.New(fmt.Sprintf("No prices found for item %s. Error: %s", itemId, iterationError.Error()))
	}; if len(prices.Prices) == 0 {
		return request.Prices{}, errors.New(fmt.Sprintf("No prices found for item %s", itemId))
	}
	return prices, nil
}

func AddPrices(prices request.Prices, db *sql.DB) error {
	itemId := prices.Item
	stmt, err := db.Prepare(`INSERT INTO price(item,price,date) VALUES ($1,$2,$3)`); if err != nil {
		return err
	}
	defer stmt.Close()
	for _, price := range prices.Prices {
		_, insertError := stmt.Exec(itemId, price.Price, price.Date)
		if insertError != nil {
			return insertError
		}
	}
	return nil
}

func EditPrice(prices request.Prices, db *sql.DB) error {
	deleteError := DeletePrice(prices.Item, db); if deleteError != nil {
		return deleteError
	}
	stmtInsert, err := db.Prepare(`INSERT INTO price(item,price,date) VALUES ($1,$2,$3)`); if err != nil {
		return err
	}
	defer stmtInsert.Close()
	for _, price := range prices.Prices {
		_, insertError := stmtInsert.Exec(prices.Item, price.Price, price.Date)
		if insertError != nil {
			return insertError
		}
	}
	return nil
}

func DeletePrice(itemId string, db *sql.DB) error {
	stmtDelete, err := db.Prepare(`DELETE FROM price WHERE item=$1`); if err != nil {
		return err
	}
	defer stmtDelete.Close()
	_, deleteError := stmtDelete.Exec(itemId); if deleteError != nil {
		return deleteError
	}
	return nil
}
