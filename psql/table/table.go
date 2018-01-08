package table

var ORDER = [9]string {
	"manufacturer",
	"currency",
	"trend",
	"item",
	"price",
	"review",
	"category",
	"forecast",
	"category_item",
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
	"Category_Item": "category_item",
}

var CREATETABLES = map[string]string{
	"manufacturer":
	"CREATE TABLE IF NOT EXISTS " + TABLES["Manufacturer"] + " (" +
		"name text PRIMARY KEY, " +
		"id smallint NOT NULL DEFAULT nextval('manufacturer_id_seq')" +
		")",
	"currency":
	"CREATE TABLE IF NOT EXISTS " + TABLES["Currency"] + " (" +
		"name currency_name NOT NULL, " +
		"date text NOT NULL, " +
		"value double precision NOT NULL," +
		"PRIMARY KEY (name,date,value)" +
		")",
	"trend":
	"CREATE TABLE IF NOT EXISTS " + TABLES["Trend"] + " (" +
		"manufacturer text REFERENCES " + TABLES["Manufacturer"] + "(name) ON DELETE CASCADE ON UPDATE CASCADE, " +
		"date text, " +
		"value double precision NOT NULL, " +
		"PRIMARY KEY(manufacturer, date, value)" +
		")",
	"item":
	"CREATE TABLE IF NOT EXISTS " + TABLES["Item"] + " (" +
		"item text PRIMARY KEY, " +
		"manufacturer text REFERENCES " + TABLES["Manufacturer"] + "(name) ON DELETE CASCADE ON UPDATE CASCADE DEFAULT 'no_manufacturer', " +
		"url text DEFAULT '', " +
		"image text DEFAULT '', " +
		"title text DEFAULT '', " +
		"description text DEFAULT '', " +
		"id smallint NOT NULL DEFAULT nextval('item_id_seq'), " +
		"has_reviews boolean DEFAULT false" +
		")",
	"price":
	"CREATE TABLE IF NOT EXISTS " + TABLES["Price"] + " (" +
		"id smallint NOT NULL DEFAULT nextval('price_id_seq'), " +
		"item text REFERENCES " + TABLES["Item"] + "(item) ON DELETE CASCADE ON UPDATE CASCADE, " +
		"price double precision NOT NULL, " +
		"date text NOT NULL, " +
		"flag boolean DEFAULT false, " +
		"PRIMARY KEY (item,price,date)" +
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
		"category text PRIMARY KEY" +
		")",
	"forecast":
	"CREATE TABLE IF NOT EXISTS " + TABLES["Forecast"] + " (" +
		"id smallint NOT NULL DEFAULT nextval('forecast_id_seq'), " +
		"name text, " +
		"item text REFERENCES " + TABLES["Item"] + "(item) ON DELETE CASCADE ON UPDATE CASCADE, " +
		"price double precision NOT NULL, " +
		"date text NOT NULL," +
		"PRIMARY KEY (name,item,price,date)" +
		")",
	"category_item":
	"CREATE TABLE IF NOT EXISTS " + TABLES["Category_Item"] + " (" +
		"item text REFERENCES item (item) ON UPDATE CASCADE ON DELETE CASCADE, " +
		"category text REFERENCES category (category) ON UPDATE CASCADE ON DELETE CASCADE," +
		"CONSTRAINT category_item_pk PRIMARY KEY (item, category)" +
		")",
}
