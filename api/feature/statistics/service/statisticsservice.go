package service

import (
	"database/sql"
	"github.com/andream16/go-storm/model/request"
	"fmt"
	"errors"
)

func DeleteStatistics(itemId, testSize string, db *sql.DB) error {
	stmtDelete, err := db.Prepare(`DELETE FROM statistics WHERE item=$1 AND test_size=$2`); if err != nil {
		defer stmtDelete.Close()
		return err
	}
	defer stmtDelete.Close()
	_, deleteError := stmtDelete.Exec(itemId, testSize); if deleteError != nil {
		return deleteError
	}
	return nil
}

func AddStatistics(statistics request.Statistics, db *sql.DB) error {
	itemId := statistics.Item
	testSize := statistics.TestSize
	statisticsEntries, _ := GetStatisticsByItemAndForecastTestSize(itemId, testSize, db); if len(statisticsEntries.Forecast) > 0 {
		deleteStatisticsError := DeleteStatistics(itemId, testSize, db); if deleteStatisticsError != nil {
			return deleteStatisticsError
		}
	}
	stmt, err := db.Prepare(`INSERT INTO statistics(name,item,price,date,score,test_size) VALUES ($1,$2,$3,$4,$5,$6)`); if err != nil {
		defer stmt.Close()
		return err
	}
	defer stmt.Close()
	statisticsType := statistics.Name
	for _, forecast := range statistics.Forecast {
		_, insertError := stmt.Exec(statisticsType, itemId, forecast.Price, forecast.Date, forecast.Score, forecast.TestSize)
		if insertError != nil {
			return insertError
		}
	}
	return nil
}

func GetStatisticsByItemAndForecastTestSize(itemId, testSize string, db *sql.DB) (request.Statistics, error) {
	emptyResponse := request.Statistics{Item: itemId, Name: "", Forecast: []request.StatisticsEntry{}}
	stmt, err := db.Prepare(`SELECT price,date,score,test_size FROM statistics WHERE item=$1 AND test_size=$2 ORDER BY date ASC`); if err != nil {
		defer stmt.Close()
		return request.Statistics{}, err
	}
	defer stmt.Close()
	rows, queryError := stmt.Query(itemId, testSize); if queryError != nil {
		return emptyResponse, nil
	}
	defer rows.Close()
	statisticsNameQuery, forecastNameQueryError := db.Prepare(`SELECT name FROM statistics WHERE item=$1 LIMIT 1`)
	if forecastNameQueryError != nil {
		return emptyResponse, nil
	}
	defer statisticsNameQuery.Close()
	var statisticsName string
	nameError := statisticsNameQuery.QueryRow(itemId).Scan(&statisticsName); if nameError != nil || len(statisticsName) == 0 {
		return emptyResponse, nil
	}
	var statistics request.Statistics
	statistics.Name = statisticsName
	statistics.Item = itemId
	scoreAndTestSizeHaveBeenSet := false
	for rows.Next() {
		var statistic request.StatisticsEntry
		rowError := rows.Scan(&statistic.Price, &statistic.Date, &statistic.Score, &statistic.TestSize); if rowError != nil {
			return request.Statistics{}, errors.New(fmt.Sprintf("Unable to unmarshal statistics for item %s. Error: %s", itemId, rowError.Error()))
		}
		if !scoreAndTestSizeHaveBeenSet {
			statistics.Score = statistic.Score
			statistics.TestSize = statistic.TestSize
			scoreAndTestSizeHaveBeenSet = true
		}
		statistics.Forecast = append(statistics.Forecast, request.StatisticsEntry{
			Price: statistic.Price,
			Date: statistic.Date,
		})
	}
	iterationError := rows.Err(); if iterationError != nil {
		return request.Statistics{}, errors.New(fmt.Sprintf("No statistics found for item %s. Error: %s", itemId, iterationError.Error()))
	}
	return statistics, nil
}
