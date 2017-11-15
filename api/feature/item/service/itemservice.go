package service

import (
	"database/sql"
	"github.com/andream16/go-storm/model/request"
	"github.com/go-errors/errors"
	"fmt"
)

func GetItem(itemId string, db *sql.DB) request.Item {
	var item request.Item
	db.QueryRow(`SELECT * FROM item WHERE item=$1`, itemId).
		Scan(&item.Item, &item.Manufacturer, &item.URL, &item.Image, &item.Title, &item.Description, &item.HasReviews)
	return item
}

func GetItems(page, size int, db *sql.DB) ([]request.Item, error) {
	var items []request.Item
	var start, end int
	if page == 1 {
		start = 1; end = size
	} else {
		start = ((page -1) * size) + 1; end = page * size
	}
	rows, queryError := db.Query(`SELECT item,manufacturer,url,image,title,description,has_reviews FROM item WHERE id BETWEEN $1 AND $2`, start, end); if queryError != nil {
		return []request.Item{}, errors.New(fmt.Sprintf("Unable to get items for page %v and size %v. Error: %s", page, size, queryError.Error()))
	}
	defer rows.Close()
	for rows.Next() {
		var item request.Item
		rowError := rows.Scan(&item.Item, &item.Manufacturer, &item.URL, &item.Image, &item.Title, &item.Description, &item.HasReviews)
		if rowError != nil {
			return []request.Item{}, errors.New(fmt.Sprintf("Unable to unmarshal items for page %v and size %v. Error: %s", page, size, rowError.Error()))
		}
		items = append(items, item)
	}; iterationError := rows.Err(); if iterationError != nil {
		return []request.Item{}, errors.New(fmt.Sprintf("No items found for page %v and size %v. Error: %s", page, size, iterationError.Error()))
	}; if len(items) == 0 {
		return []request.Item{}, errors.New(fmt.Sprintf("No items found for page %v and size %v", page, size))
	}
	return items, nil
}

func AddItem(item request.Item, db *sql.DB) error {
	_, insertError := db.Query(`INSERT INTO item(item,manufacturer,url,image,title,description,has_reviews)` +
										  ` VALUES($1,$2,$3,$4,$5,$6,$7) ON CONFLICT (item) DO UPDATE SET` +
										  ` manufacturer=$2, url=$3, image=$4, title=$5, description=$6, has_reviews=$7`,
		&item.Item, &item.Manufacturer, &item.URL, &item.Image, &item.Title, &item.Description, &item.HasReviews)
	if insertError != nil {
		return insertError
	}
	return nil
}

func EditItem(item request.Item, db *sql.DB) error {
	_, updateError := db.Query(`UPDATE item SET` +
										` manufacturer=$1, url=$2, image=$3, title=$4, description=$5, has_reviews=$6` +
	                                    ` WHERE item = $7`,
		&item.Manufacturer, &item.URL, &item.Image, &item.Title, &item.Description, &item.HasReviews, &item.Item)
	if updateError != nil {
		return updateError
	}
	return nil
}

func DeleteItem(itemId string, db *sql.DB) error {
	_, deleteError := db.Query(`DELETE FROM item WHERE item=$1`, itemId); if deleteError != nil {
		return deleteError
	}
	return nil
}