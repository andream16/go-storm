package service

import (
	"database/sql"
	"github.com/andream16/go-storm/model/request"
)

func AddTmpReview(review request.ReviewTmp, db *sql.DB) error {
	deleteError := deleteTmpReviewsByItem(review.Item, db); if deleteError != nil {
		return deleteError
	}
	stmtInsert, err := db.Prepare(`INSERT INTO reviewtmp(item,content,sentiment,stars,date) VALUES ($1,$2,$3,$4,$5)`); if err != nil {
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

func deleteTmpReviewsByItem(item string, db *sql.DB) error {
	stmtDelete, err := db.Prepare(`DELETE FROM reviewtmp WHERE item= $1`); if err != nil {
		return err
	}
	defer stmtDelete.Close()
	_, deleteError := stmtDelete.Exec(item); if deleteError != nil {
		return deleteError
	}
	return nil
}
