package service

import (
	"database/sql"
	"github.com/andream16/go-storm/model/request"
)

func AddTmpReview(review request.ReviewTmp, db *sql.DB) error {
	stmtInsert, err := db.Prepare(`INSERT INTO review_tmp(item,content,sentiment,stars,date) VALUES ($1,$2,$3,$4,$5)`); if err != nil {
		return err
	}
	defer stmtInsert.Close()
	for _, reviewEntry := range review.Reviews {
		_, insertError := stmtInsert.Exec(review.Item, reviewEntry.Content, reviewEntry.Sentiment, reviewEntry.Stars, reviewEntry.Date)
		if insertError != nil {
			return insertError
		}
	}
	return nil
}