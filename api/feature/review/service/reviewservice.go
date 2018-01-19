package service

import (
	"database/sql"
	"github.com/andream16/go-storm/model/request"
	"fmt"
	"errors"
)

func GetReviewByItem(item string, db *sql.DB) (request.Review, error) {
	stmt, err := db.Prepare(`SELECT item,date,content,sentiment,stars FROM review WHERE item=$1 ORDER BY date ASC`); if err != nil {
		return request.Review{}, err
	}
	defer stmt.Close()
	rows, getReviewError := stmt.Query(item); if getReviewError != nil {
		return request.Review{}, errors.New(fmt.Sprintf("Unable to find review rows for item %s", item))
	}
	defer rows.Close()
	var review request.Review; review.Item = item
	for rows.Next() {
		var reviewEntry request.ReviewEntry
		rowError := rows.Scan(&reviewEntry.Item, &reviewEntry.Date, &reviewEntry.Content, &reviewEntry.Sentiment, &reviewEntry.Stars); if rowError != nil {
			return request.Review{}, errors.New(fmt.Sprintf("Unable to unmarshal review entries for item %s. Error: %s", item, rowError.Error()))
		}
		review.Reviews = append(review.Reviews, reviewEntry)
	}; iterationError := rows.Err(); if iterationError != nil {
		return request.Review{}, errors.New(fmt.Sprintf("No reviews entries found for item %s. Error: %s", item, iterationError.Error()))
	}; if len(review.Reviews) == 0 {
		return request.Review{}, errors.New(fmt.Sprintf("No reviews entries found for item %s", item))
	}
	return review, nil
}


func AddReviews(review request.Review, db *sql.DB) error {
	deleteError := DeleteReviewsByItem(review.Item, db); if deleteError != nil {
		return deleteError
	}
	stmtInsert, err := db.Prepare(`INSERT INTO review(item,content,sentiment,stars,date) VALUES ($1,$2,$3,$4,$5)`); if err != nil {
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

func EditReviewByItem(review request.Review, db *sql.DB) error {
	deleteError := DeleteReviewsByItem(review.Item, db); if deleteError != nil {
		return deleteError
	}
	return  AddReviews(review, db)
}

func DeleteReviewsByItem(item string, db *sql.DB) error {
	stmtDelete, err := db.Prepare(`DELETE FROM review WHERE item= $1`); if err != nil {
		return err
	}
	defer stmtDelete.Close()
	_, deleteError := stmtDelete.Exec(item); if deleteError != nil {
		return deleteError
	}
	return nil
}
