package service

import (
	"github.com/andream16/go-storm/model/request"
	"database/sql"
	"fmt"
	"errors"
)

func GetForecasts(itemId string, db *sql.DB) (request.Forecast, error) {
	stmt, err := db.Prepare(`SELECT price,date FROM forecast WHERE item=$1 ORDER BY date ASC`); if err != nil {
		defer stmt.Close()
		return request.Forecast{}, err
	}
	defer stmt.Close()
	rows, queryError := stmt.Query(itemId); if queryError != nil {
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
	stmt, err := db.Prepare(`INSERT INTO forecast(name,item,price,date) VALUES ($1,$2,$3,$4)`); if err != nil {
		defer stmt.Close()
		return err
	}
	defer stmt.Close()
	forecastType := forecasts.Name
	for _, forecast := range forecasts.Forecast {
		_, insertError := stmt.Exec(forecastType, itemId, forecast.Price, forecast.Date)
		if insertError != nil {
			return insertError
		}
	}
	return nil
}

func EditForecast(forecasts request.Forecast, db *sql.DB) error {
	stmtInsert, err := db.Prepare(`INSERT INTO forecast(name,item,price,date) VALUES ($1,$2,$3,$4)`); if err != nil {
		defer stmtInsert.Close()
		return err
	}
	defer stmtInsert.Close()
	forecastType := forecasts.Name
	deleteError := DeleteForecast(forecasts.Item, db); if deleteError != nil {
		return deleteError
	}
	for _, forecast := range forecasts.Forecast {
		_, insertError := stmtInsert.Exec(forecastType, forecasts.Item, forecast.Price, forecast.Date)
		if insertError != nil {
			return insertError
		}
	}
	return nil
}

func DeleteForecast(itemId string, db *sql.DB) error {
	stmtDelete, err := db.Prepare(`DELETE FROM forecast WHERE item=$1`); if err != nil {
		defer stmtDelete.Close()
		return err
	}
	defer stmtDelete.Close()
	_, deleteError := stmtDelete.Exec(itemId); if deleteError != nil {
		return deleteError
	}
	return nil
}
