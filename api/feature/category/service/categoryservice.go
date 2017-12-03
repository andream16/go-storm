package service

import (
	"database/sql"
	"github.com/andream16/go-storm/model/request"
	"fmt"
	"errors"
)

func GetItemsByCategory(category string, db *sql.DB) ([]request.Item, error) {
	stmt, err := db.Prepare(`SELECT item FROM category_item WHERE category=$1 ORDER BY item ASC`); if err != nil {
		defer stmt.Close()
		return []request.Item{}, err
	}
	defer stmt.Close()
	rows, queryError := stmt.Query(category); if queryError != nil {
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
	stmt, err := db.Prepare(`SELECT category FROM category_item WHERE item=$1 ORDER BY category ASC`); if err != nil {
		defer stmt.Close()
		return []request.Category{}, err
	}
	defer stmt.Close()
	rows, queryError := stmt.Query(itemId); if queryError != nil {
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
	stmtCategory, err := db.Prepare(`INSERT INTO category(category) VALUES($1) ON CONFLICT (category) DO UPDATE SET category=$1`); if err != nil {
		defer stmtCategory.Close()
		return err
	}
	defer stmtCategory.Close()
	stmtCategoryItems, err := db.Prepare(`INSERT INTO category_item(category,item) VALUES($1,$2)`); if err != nil {
		defer stmtCategoryItems.Close()
		return err
	}
	defer stmtCategoryItems.Close()
	for _, category := range categoriesByItem.Categories {
	_, insertCategoriesError := stmtCategory.Exec(&category); if insertCategoriesError != nil {
			return insertCategoriesError
		}
	_, insertCategoriesItemsError := stmtCategoryItems.Exec(&category, &categoriesByItem.Item); if insertCategoriesItemsError != nil {
			return insertCategoriesItemsError
		}
	}
	return nil
}

func EditCategory(categoriesByItem request.CategoryRequest, db *sql.DB) error  {
	stmt, err := db.Prepare(`DELETE FROM category_item WHERE item=$1`); if err != nil {
		defer stmt.Close()
		return err
	}
	defer stmt.Close()
	_, updateCategoriesError := stmt.Exec(&categoriesByItem.Item); if updateCategoriesError != nil {
		return updateCategoriesError
	}
	return AddCategoriesByItem(categoriesByItem, db)
}

func DeleteCategory(categoriesByItem request.CategoryRequest, db *sql.DB) error  {
	stmt, err := db.Prepare(`DELETE FROM category_item WHERE item=$1 AND category=$2`); if err != nil {
		defer stmt.Close()
		return err
	}
	defer stmt.Close()
	for _, category := range categoriesByItem.Categories {
		_, deleteCategoriesError := stmt.Exec(&categoriesByItem.Item, &category); if deleteCategoriesError != nil {
			return deleteCategoriesError
		}
	}
	return nil
}
