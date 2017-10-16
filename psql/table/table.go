package table

var ORDER = [8]string {
	"manufacturer",
	"currency",
	"trend",
	"item",
	"price",
	"review",
	"category",
	"forecast",
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
	"manufacturer":
	"CREATE TABLE IF NOT EXISTS " + TABLES["Manufacturer"] + " (" +
		"name text PRIMARY KEY, " +
		"id smallint NOT NULL DEFAULT nextval('manufacturer_id_seq')" +
		")",
	"currency":
	"CREATE TABLE IF NOT EXISTS " + TABLES["Currency"] + " (" +
		"name text PRIMARY KEY, " +
		"date text, " +
		"value numeric(10,2)" +
		")",
	"trend":
	"CREATE TABLE IF NOT EXISTS " + TABLES["Trend"] + " (" +
		"id smallint NOT NULL DEFAULT nextval('trend_id_seq') PRIMARY KEY, " +
		"manufacturer text REFERENCES " + TABLES["Manufacturer"] + "(name) ON DELETE CASCADE ON UPDATE CASCADE, " +
		"date text, " +
		"value numeric(10,2)" +
		")",
	"item":
	"CREATE TABLE IF NOT EXISTS " + TABLES["Item"] + " (" +
		"item text PRIMARY KEY, " +
		"manufacturer text REFERENCES " + TABLES["Manufacturer"] + "(name) ON DELETE CASCADE ON UPDATE CASCADE, " +
		"url text, " +
		"image text, " +
		"title text, " +
		"description text, " +
		"id smallint NOT NULL DEFAULT nextval('item_id_seq'), " +
		"has_reviews boolean DEFAULT false" +
		")",
	"price":
	"CREATE TABLE IF NOT EXISTS " + TABLES["Price"] + " (" +
		"id smallint NOT NULL DEFAULT nextval('price_id_seq') PRIMARY KEY, " +
		"item text REFERENCES " + TABLES["Item"] + "(item) ON DELETE CASCADE ON UPDATE CASCADE, " +
		"price numeric(10,2), " +
		"date text, " +
		"flag boolean DEFAULT false" +
		")",
	"review":
	"CREATE TABLE IF NOT EXISTS " + TABLES["Review"] + " (" +
		"id smallint NOT NULL DEFAULT nextval('review_id_seq') PRIMARY KEY, " +
		"item text REFERENCES " + TABLES["Item"] + "(item) ON DELETE CASCADE ON UPDATE CASCADE, " +
		"content text, " +
		"date text, " +
		"sentiment smallint, constraint valid_sentiment check(sentiment BETWEEN 0 AND 1), " +
		"stars smallint, constraint valid_stars check(stars BETWEEN 1 AND 5)" +
		")",
	"category":
	"CREATE TABLE IF NOT EXISTS " + TABLES["Category"] + " (" +
		"name text PRIMARY KEY, " +
		"item text REFERENCES " + TABLES["Item"] + "(item) ON DELETE CASCADE ON UPDATE CASCADE" +
		")",
	"forecast":
	"CREATE TABLE IF NOT EXISTS " + TABLES["Forecast"] + " (" +
		"id smallint NOT NULL DEFAULT nextval('forecast_id_seq') PRIMARY KEY, " +
		"name text, " +
		"item text REFERENCES " + TABLES["Item"] + "(item) ON DELETE CASCADE ON UPDATE CASCADE, " +
		"price numeric(10,2), " +
		"date text" +
		")",
}
