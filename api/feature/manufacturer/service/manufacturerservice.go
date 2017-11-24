package service

import (
	"database/sql"
	"github.com/andream16/go-storm/model/request"
	"fmt"
	"errors"
)

func GetItemsByManufacturer(manufacturer string, db *sql.DB) ([]request.Item, error) {
	rows, queryError := db.Query(`SELECT item FROM item WHERE manufacturer=$1 ORDER BY item ASC`, manufacturer); if queryError != nil {
		return []request.Item{}, errors.New(fmt.Sprintf("Unable to get item for manufacturer %s. Error: %s", manufacturer, queryError.Error()))
	}
	defer rows.Close()
	var items []request.Item
	for rows.Next() {
		var item request.Item
		rowError := rows.Scan(&item.Item); if rowError != nil {
			return []request.Item{}, errors.New(fmt.Sprintf("Unable to unmarshal items for manufacturer %s. Error: %s", manufacturer, rowError.Error()))
		}
		items = append(items, item)
	}; iterationError := rows.Err(); if iterationError != nil {
		return []request.Item{}, errors.New(fmt.Sprintf("No items found for manufacturer %s. Error: %s", manufacturer, iterationError.Error()))
	}; if len(items) == 0 {
		return []request.Item{}, errors.New(fmt.Sprintf("No items found for manufacturer %s", manufacturer))
	}
	return items, nil
}

func GetManufacturerByItem(itemId string, db *sql.DB) (request.Manufacturer, error) {
	rows, queryError := db.Query(`SELECT manufacturer FROM item WHERE item=$1`, itemId); if queryError != nil {
		return request.Manufacturer{}, errors.New(fmt.Sprintf("Unable to get manufacturer for item %s. Error: %s", itemId, queryError.Error()))
	}
	defer rows.Close()
	var manufacturer request.Manufacturer
	for rows.Next() {
		rowError := rows.Scan(&manufacturer.Name)
		if rowError != nil {
			return request.Manufacturer{}, errors.New(fmt.Sprintf("Unable to unmarshal manufacturer for item %s. Error: %s", itemId, rowError.Error()))
		}
	}

	return manufacturer, nil
}

func AddManufacturer(manufacturer request.Manufacturer, db *sql.DB) error {
	_, insertManufacturerError := db.Query(`INSERT INTO manufacturer(name)` +
			` VALUES($1)`,
			&manufacturer.Name); if insertManufacturerError != nil {
			return insertManufacturerError
		}

	return nil
}

func DeleteManufacturer(manufacturer request.Manufacturer, db *sql.DB) error  {
		_, deleteManufacturerError := db.Query(`DELETE FROM manufacturer WHERE name=$1`, &manufacturer.Name); if deleteManufacturerError != nil {
			return deleteManufacturerError
		}

	return nil
}
