package service

import (
	"database/sql"
	"github.com/andream16/go-storm/model/request"
)

func GetItem(itemId string, db *sql.DB) request.Item {
	var item request.Item
	db.QueryRow(`SELECT * FROM item WHERE item=$1`, itemId).
		Scan(&item.Item, &item.Manufacturer, &item.URL, &item.Image, &item.Title, &item.Description, &item.HasReviews)
	return item
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