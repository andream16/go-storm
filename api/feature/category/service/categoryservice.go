package service

import (
	"database/sql"
	"github.com/andream16/go-storm/model/request"
	"fmt"
	"errors"
)

func GetItemsByCategory(category string, db *sql.DB) ([]request.Item, error) {
	rows, queryError := db.Query(`SELECT item FROM category_item WHERE category=$1 ORDER BY item ASC`, category); if queryError != nil {
		return []request.Item{}, errors.New(fmt.Sprintf("Unable to get item for category %s. Error: %s", category, queryError.Error()))
	}
	defer rows.Close()
	var items []request.Item
	for rows.Next() {
		var item request.Item
		rowError := rows.Scan(&item.Item); if rowError != nil {
			return []request.Item{}, errors.New(fmt.Sprintf("Unable to unmarshal items for category %s. Error: %s", category, rowError.Error()))
		}
		items = append(items, item)
	}; iterationError := rows.Err(); if iterationError != nil {
		return []request.Item{}, errors.New(fmt.Sprintf("No items found for category %s. Error: %s", category, iterationError.Error()))
	}; if len(items) == 0 {
		return []request.Item{}, errors.New(fmt.Sprintf("No items found for category %s", category))
	}
	return items, nil
}

func GetCategoriesByItem(itemId string, db *sql.DB) ([]request.Category, error) {
	rows, queryError := db.Query(`SELECT category FROM category_item WHERE item=$1 ORDER BY category ASC`, itemId); if queryError != nil {
		return []request.Category{}, errors.New(fmt.Sprintf("Unable to get categories for item %s. Error: %s", itemId, queryError.Error()))
	}
	defer rows.Close()
	var categories []request.Category
	for rows.Next() {
		var category request.Category
		rowError := rows.Scan(&category.Category); if rowError != nil {
			return []request.Category{}, errors.New(fmt.Sprintf("Unable to unmarshal categories for item %s. Error: %s", itemId, rowError.Error()))
		}
		categories = append(categories, category)
	}; iterationError := rows.Err(); if iterationError != nil {
		return []request.Category{}, errors.New(fmt.Sprintf("No categories found for item %s. Error: %s", itemId, iterationError.Error()))
	}; if len(categories) == 0 {
		return []request.Category{}, errors.New(fmt.Sprintf("No categories found for item %s", itemId))
	}
	return categories, nil
}

func AddCategoriesByItem(categoriesByItem request.CategoryRequest, db *sql.DB) error {
	for _, category := range categoriesByItem.Categories {
		_, insertCategoriesError := db.Query(`INSERT INTO category(category)` +
			` VALUES($1) ON CONFLICT (category) DO UPDATE SET` +
			` category=$1`,
			&category); if insertCategoriesError != nil {
				return insertCategoriesError
			}
		_, insertCategoriesItemsError := db.Query(`INSERT INTO category_item(category,item)` +
			` VALUES($1,$2)`,
			&category, &categoriesByItem.Item); if insertCategoriesItemsError != nil {
			return insertCategoriesItemsError
		}
	}
	return nil
}

func EditCategory(categoriesByItem request.CategoryRequest, db *sql.DB) error  {
	_, updateCategoriesError := db.Query(`DELETE FROM category_item WHERE item=$1`, &categoriesByItem.Item); if updateCategoriesError != nil {
		return updateCategoriesError
	}
	return AddCategoriesByItem(categoriesByItem, db)
}

func DeleteCategory(categoriesByItem request.CategoryRequest, db *sql.DB) error  {
	for _, category := range categoriesByItem.Categories {
		_, deleteCategoriesError := db.Query(`DELETE FROM category_item WHERE item=$1 AND category=$2`, &categoriesByItem.Item, &category); if deleteCategoriesError != nil {
			return deleteCategoriesError
		}
	}
	return nil
}
