package service

import (
	"github.com/andream16/go-storm/model/request"
	"database/sql"
	"fmt"
	"errors"
)

func GetPrices(itemId string, db *sql.DB) (request.Prices, error) {
	rows, queryError := db.Query(`SELECT price,date FROM price WHERE item=$1 ORDER BY date ASC`, itemId); if queryError != nil {
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
	for _, price := range prices.Prices {
		_, insertError := db.Query(`INSERT INTO price(item,price,date) VALUES ($1,$2,$3)`, itemId, price.Price, price.Date)
		if insertError != nil {
			return insertError
		}
	}
	return nil
}

func EditPrice(prices request.Prices, db *sql.DB) error {
	_, deleteError := db.Query(`DELETE FROM price WHERE item=$1`, prices.Item); if deleteError != nil {
		return deleteError
	}
	for _, price := range prices.Prices {
		_, insertError := db.Query(`INSERT INTO price(item,price,date) VALUES ($1,$2,$3)`, prices.Item, price.Price, price.Date)
		if insertError != nil {
			return insertError
		}
	}
	return nil
}

func DeletePrice(itemId string, db *sql.DB) error {
	_, deleteError := db.Query(`DELETE FROM price WHERE item=$1`, itemId); if deleteError != nil {
		return deleteError
	}
	return nil
}
