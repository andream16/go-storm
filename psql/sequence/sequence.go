package sequence

var SEQUENCES = map[string]string {
	"Manufacturer" : "CREATE SEQUENCE manufacturer_id_seq",
	"Item"		   : "CREATE SEQUENCE item_id_seq",
	"Trend"        : "CREATE SEQUENCE trend_id_seq",
	"Price"        : "CREATE SEQUENCE price_id_seq",
	"Review"       : "CREATE SEQUENCE review_id_seq",
	"Forecast"     : "CREATE SEQUENCE forecast_id_seq",
}
