package endpoint

import (
	"net/http"
	"fmt"
	"strconv"
	"github.com/codegangsta/negroni"
	"strings"
	"log"
	"github.com/rs/cors"
	"github.com/andream16/go-storm/configuration"
	"database/sql"
	pingHandler "github.com/andream16/go-storm/api/ping/handler"
	itemHandler "github.com/andream16/go-storm/api/item/handler"
	priceHandler "github.com/andream16/go-storm/api/price/handler"
	manufacturerHandler "github.com/andream16/go-storm/api/manufacturer/handler"
	reviewHandler "github.com/andream16/go-storm/api/review/handler"
	forecastHandler "github.com/andream16/go-storm/api/forecast/handler"
	categoryHandler "github.com/andream16/go-storm/api/category/handler"
	currencyHandler "github.com/andream16/go-storm/api/currency/handler"
	trendHandler "github.com/andream16/go-storm/api/trend/handler"
)

func InitializeEndpoint(conf *configuration.Configuration, db *sql.DB) {
	fmt.Println("Initializing endpoints ...")
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET, POST, PUT, DELETE"},
	})
	mux := http.NewServeMux()
	mux.HandleFunc("/", 			    pingHandler.PingHandler)
	mux.HandleFunc("/item", 			itemHandler.ItemHandler(db))
	mux.HandleFunc("/price", 			priceHandler.PriceHandler(db))
	mux.HandleFunc("/manufacturer", 	manufacturerHandler.ManufacturerHandler(db))
	mux.HandleFunc("/review", 		reviewHandler.ReviewHandler(db))
	mux.HandleFunc("/forecast", 		forecastHandler.ForecastHandler(db))
	mux.HandleFunc("/category", 		categoryHandler.CategoryHandler(db))
	mux.HandleFunc("/currency", 		currencyHandler.CurrencyHandler(db))
	mux.HandleFunc("/trend", 			trendHandler.TrendHandler(db))
	port := strconv.Itoa(conf.Server.Port)
	n := negroni.Classic(); n.Use(c); n.UseHandler(mux)
	fmt.Println("Started serversharer at port :" + port + ". Now listening . . .")
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(strings.Join([]string{":", port}, ""), n))
}