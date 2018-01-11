package service

import (
	"database/sql"
	"github.com/andream16/go-storm/model/request"
	"github.com/go-errors/errors"
	"fmt"
)

func GetItem(key string, value string, db *sql.DB) request.Item {
	var item request.Item
	stmt, err := db.Prepare(fmt.Sprintf(`SELECT item,manufacturer,url,image,title,description,has_reviews FROM item WHERE %s = $1`, key)); if err != nil {
		return request.Item{}
	}
	defer stmt.Close()
	selectError := stmt.QueryRow(value).Scan(&item.Item, &item.Manufacturer, &item.URL, &item.Image, &item.Title, &item.Description, &item.HasReviews)
	if selectError != nil {
		return request.Item{}
	}
	return item
}

func GetItems(page, size int, db *sql.DB) (request.Items, error) {
	var items []request.Item
	var start, end int
	hasNext := hasNextPaginatedItems(page, size, db)
	stmt, err := db.Prepare(`SELECT item,manufacturer,url,image,title,description,has_reviews FROM item WHERE id BETWEEN $1 AND $2`); if err != nil {
		return request.Items{}, err
	}
	defer stmt.Close()
	if page == 1 {
		start = 1; end = size
	} else {
		start = ((page - 1) * size) + 1; end = page * size
	}
	rows, queryError := stmt.Query(start, end); if queryError != nil {
		return request.Items{}, errors.New(fmt.Sprintf("Unable to get items for page %v and size %v. Error: %s", page, size, queryError.Error()))
	}
	defer rows.Close()
	for rows.Next() {
		var item request.Item
		rowError := rows.Scan(&item.Item, &item.Manufacturer, &item.URL, &item.Image, &item.Title, &item.Description, &item.HasReviews)
		if rowError != nil {
			return request.Items{}, errors.New(fmt.Sprintf("Unable to unmarshal items for page %v and size %v. Error: %s", page, size, rowError.Error()))
		}
		items = append(items, item)
	}; iterationError := rows.Err(); if iterationError != nil {
		return request.Items{}, errors.New(fmt.Sprintf("No items found for page %v and size %v. Error: %s", page, size, iterationError.Error()))
	}; if len(items) == 0 {
		return request.Items{}, errors.New(fmt.Sprintf("No items found for page %v and size %v", page, size))
	}
	return request.Items{Items: items, HasNext: hasNext}, nil
}

func hasNextPaginatedItems(page, size int, db *sql.DB) bool {
	var count int
	stmt, err := db.Prepare(`SELECT COUNT(item) FROM item`); if err != nil {
		return false
	}
	defer stmt.Close()
	selectError := stmt.QueryRow().Scan(&count); if selectError != nil {
		return false
	}
	if count > (page * size) {
		return true
	}
	return false
}

func AddItem(item request.Item, db *sql.DB) error {
	stmt, err := db.Prepare(`INSERT INTO item(item,manufacturer,url,image,title,description,has_reviews) VALUES($1,$2,$3,$4,$5,$6,$7) ON CONFLICT (item) DO UPDATE SET manufacturer=$2, url=$3, image=$4, title=$5, description=$6, has_reviews=$7`); if err != nil {
		return err
	}
	defer stmt.Close()
	_, insertError := stmt.Exec(&item.Item, &item.Manufacturer, &item.URL, &item.Image, &item.Title, &item.Description, &item.HasReviews)
	if insertError != nil {
		return insertError
	}
	return nil
}

func EditItem(item request.Item, db *sql.DB) error {
	stmt, err := db.Prepare(`UPDATE item SET manufacturer=$1, url=$2, image=$3, title=$4, description=$5, has_reviews=$6 WHERE item = $7`); if err != nil {
		return err
	}
	defer stmt.Close()
	_, updateError := stmt.Exec(&item.Manufacturer, &item.URL, &item.Image, &item.Title, &item.Description, &item.HasReviews, &item.Item)
	if updateError != nil {
		return updateError
	}
	return nil
}

func DeleteItem(itemId string, db *sql.DB) error {
	stmt, err := db.Prepare(`DELETE FROM item WHERE item=$1`); if err != nil {
		return err
	}
	defer stmt.Close()
	_, deleteError := stmt.Exec(itemId)
	if deleteError != nil {
		return deleteError
	}
	return nil
}