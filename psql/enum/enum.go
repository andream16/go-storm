package enum

var ENUMERATIONS= map[string]string {
	"Forecast": "CREATE TYPE forecast_name AS ENUM ('ATCR','ATR','ACR','ATC','AT','AC','AR')",
	"Currency":"CREATE TYPE currency_name AS ENUM ('EURO', 'DOLLAR')",
}
