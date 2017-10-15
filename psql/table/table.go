package table

var ORDER = [8]string {
	"Manufacturer",
	"Currency",
	"Trend",
	"Item",
	"Price",
	"Review",
	"Category",
	"Forecast",
}

var TABLES = map[string]string {
	"Manufacturer" : "manufacturer",
	"Currency"     : "currency",
	"Trend"        : "trend",
	"Item" 		   : "item",
	"Price"        : "price",
	"Review"       : "review",
	"Category"     : "category",
	"Forecast" 	   : "forecast",
}

var CREATETABLES = map[string]string{
	"Manufacturer":
	"CREATE TABLE IF NOT EXISTS " + TABLES["Manufacturer"] + " (" +
		"name text PRIMARY KEY, " +
		"id smallint NOT NULL DEFAULT nextval('manufacturer_id_seq')" +
		")",
	"Currency":
	"CREATE TABLE IF NOT EXISTS " + TABLES["Currency"] + " (" +
		"name text PRIMARY KEY, " +
		"date text, " +
		"value numeric(10,2)" +
		")",
	"Trend":
	"CREATE TABLE IF NOT EXISTS " + TABLES["Trend"] + " (" +
		"id smallint NOT NULL DEFAULT nextval('trend_id_seq') PRIMARY KEY, " +
		"manufacturer text REFERENCES " + TABLES["Manufacturer"] + "(name) ON DELETE CASCADE ON UPDATE CASCADE, " +
		"date text, " +
		"value numeric(10,2)" +
		")",
	"Item":
	"CREATE TABLE IF NOT EXISTS " + TABLES["Item"] + " (" +
		"item text PRIMARY KEY, " +
		"manufacturer text REFERENCES " + TABLES["Manufacturer"] + "(name) ON DELETE CASCADE ON UPDATE CASCADE, " +
		"url text, " +
		"image text, " +
		"title text, " +
		"description text, " +
		"has_reviews boolean DEFAULT false" +
		")",
	"Price":
	"CREATE TABLE IF NOT EXISTS " + TABLES["Price"] + " (" +
		"id smallint NOT NULL DEFAULT nextval('price_id_seq') PRIMARY KEY, " +
		"item text REFERENCES " + TABLES["Item"] + "(item) ON DELETE CASCADE ON UPDATE CASCADE, " +
		"price numeric(10,2), " +
		"date text, " +
		"flag boolean DEFAULT false" +
		")",
	"Review":
	"CREATE TABLE IF NOT EXISTS " + TABLES["Review"] + " (" +
		"id smallint NOT NULL DEFAULT nextval('review_id_seq') PRIMARY KEY, " +
		"item text REFERENCES " + TABLES["Item"] + "(item) ON DELETE CASCADE ON UPDATE CASCADE, " +
		"content text, " +
		"date text, " +
		"sentiment smallint, constraint valid_sentiment check(sentiment => 0 and <= 1), " +
		"stars smallint, constraint valid_stars check(stars > 0 and stars < 6)" +
		")",
	"Category":
	"CREATE TABLE IF NOT EXISTS " + TABLES["Category"] + " (" +
		"name text PRIMARY KEY, " +
		"item text REFERENCES " + TABLES["Item"] + "(item) ON DELETE CASCADE ON UPDATE CASCADE" +
		")",
	"Forecast":
	"CREATE TABLE IF NOT EXISTS " + TABLES["Forecast"] + " (" +
		"id smallint NOT NULL DEFAULT nextval('forecast_id_seq') PRIMARY KEY, " +
		"name text, " +
		"item text REFERENCES " + TABLES["Item"] + "(item) ON DELETE CASCADE ON UPDATE CASCADE, " +
		"price numeric(10,2), " +
		"date text" +
		")",
}
