package service

import (
	"github.com/andream16/go-storm/model/request"
	"database/sql"
	"fmt"
	"errors"
)

func GetForecasts(itemId string, db *sql.DB) (request.Forecast, error) {
	rows, queryError := db.Query(`SELECT price,date FROM forecast WHERE item=$1 ORDER BY date ASC`, itemId); if queryError != nil {
		return request.Forecast{}, errors.New(fmt.Sprintf("Unable to get forecasts for item %s. Error: %s", itemId, queryError.Error()))
	}
	defer rows.Close()
	var forecasts request.Forecast
	forecasts.Item = itemId
		for rows.Next() {
			var forecast request.ForecastEntry
			rowError := rows.Scan(&forecast.Price, &forecast.Date); if rowError != nil {
			return request.Forecast{}, errors.New(fmt.Sprintf("Unable to unmarshal forecasts for item %s. Error: %s", itemId, rowError.Error()))
			}
			forecasts.Forecast = append(forecasts.Forecast, forecast)
		};
		iterationError := rows.Err();
		if iterationError != nil {
			return request.Forecast{}, errors.New(fmt.Sprintf("No forecasts found for item %s. Error: %s", itemId, iterationError.Error()))
		}; if len(forecasts.Forecast) == 0 {
			return request.Forecast{}, errors.New(fmt.Sprintf("No forecasts found for item %s", itemId))
			}
	return forecasts, nil
}

func AddForecasts(forecasts request.Forecast, db *sql.DB) error {
	itemId := forecasts.Item
	forecastType := forecasts.Name
	for _, forecast := range forecasts.Forecast {
		_, insertError := db.Query(`INSERT INTO forecast(name,item,price,date) VALUES ($1,$2,$3,$4)`, forecastType, itemId, forecast.Price, forecast.Date)
		if insertError != nil {
			return insertError
		}
	}
	return nil
}

func EditForecast(forecasts request.Forecast, db *sql.DB) error {
	forecastType := forecasts.Name
	_, deleteError := db.Query(`DELETE FROM forecast WHERE item=$1`, forecasts.Item); if deleteError != nil {
		return deleteError
	}
	for _, forecast := range forecasts.Forecast {
		_, insertError := db.Query(`INSERT INTO forecast(name,item,price,date) VALUES ($1,$2,$3,$4)`, forecastType, forecasts.Item, forecast.Price, forecast.Date)
		if insertError != nil {
			return insertError
		}
	}
	return nil
}

func DeleteForecast(itemId string, db *sql.DB) error {
	_, deleteError := db.Query(`DELETE FROM forecast WHERE item=$1`, itemId); if deleteError != nil {
		return deleteError
	}
	return nil
}
