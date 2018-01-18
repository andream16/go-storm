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
	pingHandler "github.com/andream16/go-storm/api/feature/ping/handler"
	itemHandler "github.com/andream16/go-storm/api/feature/item/handler"
	priceHandler "github.com/andream16/go-storm/api/feature/price/handler"
	manufacturerHandler "github.com/andream16/go-storm/api/feature/manufacturer/handler"
	reviewHandler "github.com/andream16/go-storm/api/feature/review/handler"
	reviewTmpHandler "github.com/andream16/go-storm/api/feature/review_tmp/handler"
	forecastHandler "github.com/andream16/go-storm/api/feature/forecast/handler"
	categoryHandler "github.com/andream16/go-storm/api/feature/category/handler"
	currencyHandler "github.com/andream16/go-storm/api/feature/currency/handler"
	trendHandler "github.com/andream16/go-storm/api/feature/trend/handler"
	amazonHandler "github.com/andream16/go-storm/api/feature/amazon/handler"
)

func InitializeEndpoint(conf *configuration.Configuration, db *sql.DB) {
	fmt.Println("Initializing endpoints ...")
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET, POST, PUT, DELETE"},
	})
	mux := http.NewServeMux()
	mux.HandleFunc("/api", 			    pingHandler.PingHandler)
	mux.HandleFunc("/api/item", 			itemHandler.ItemHandler(db))
	mux.HandleFunc("/api/price", 			priceHandler.PriceHandler(db))
	mux.HandleFunc("/api/manufacturer", 	manufacturerHandler.ManufacturerHandler(db))
	mux.HandleFunc("/api/review", 		reviewHandler.ReviewHandler(db))
	mux.HandleFunc("/api/review_tmp",    reviewTmpHandler.ReviewTmpHandler(db))
	mux.HandleFunc("/api/forecast", 		forecastHandler.ForecastHandler(db))
	mux.HandleFunc("/api/category", 		categoryHandler.CategoryHandler(db))
	mux.HandleFunc("/api/currency", 		currencyHandler.CurrencyHandler(db))
	mux.HandleFunc("/api/trend", 			trendHandler.TrendHandler(db))
	mux.HandleFunc("/api/amazon", 		amazonHandler.AmazonHandler(db))
	port := strconv.Itoa(conf.Server.Port)
	n := negroni.Classic(); n.Use(c); n.UseHandler(mux)
	fmt.Println("Started serversharer at port :" + port + ". Now listening . . .")
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(strings.Join([]string{":", port}, ""), n))
}