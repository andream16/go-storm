package service

import (
	"github.com/andream16/go-storm/model/request"
	"database/sql"
	manufacturerService "github.com/andream16/go-storm/api/feature/manufacturer/service"
	itemService "github.com/andream16/go-storm/api/feature/item/service"
	categoryService "github.com/andream16/go-storm/api/feature/category/service"
	reviewService "github.com/andream16/go-storm/api/feature/review/service"
)

func AddAmazonEntry(amazonEntry request.Amazon, db *sql.DB) error {
	var manufacturer = amazonEntry.Manufacturer
	var item = amazonEntry.Item
	var category = amazonEntry.Category
	var review = amazonEntry.Review
	if len(manufacturer.Manufacturer) > 0 {
		manufacturerService.AddManufacturer(request.Manufacturer{manufacturer.Manufacturer}, db)
	}
	if len(item.Item) > 0 {
		itemService.EditItem(item, db)
	}
	if len(category.Item) > 0 {
		categoryService.AddCategoriesByItem(category, db)
	}
	if len(review.Item) > 0 {
		reviewService.AddReviews(review, db)
	}
	return nil
}