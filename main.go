package main

import (
	"log"
	"net/http"

	"github.com/currency/pkg/config"
	"github.com/currency/pkg/db"
	"github.com/currency/pkg/service"
	transportHttp "github.com/currency/pkg/transport"
)

func main() {
	cfg := config.LoadConfig()

	// Connect to the database
	database, err := db.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer database.Close()

	// Initialize repository and service
	currencyRepo := db.NewPostgresCurrencyRepository(database)
	currencyService := service.NewCurrencyService(currencyRepo)

	// Set up HTTP transport
	handler := transportHttp.NewHTTPHandler(currencyService)

	// Start the HTTP server
	log.Println("Starting server on :8083")
	if err := http.ListenAndServe(":8083", handler); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
