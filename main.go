package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tobiasaagaard/localgift-api/config"
	"github.com/tobiasaagaard/localgift-api/internal/handlers"
	"github.com/tobiasaagaard/localgift-api/middleware"
)

func main() {
	cfg := config.LoadConfig()

	healthHandler := handlers.NewHealthHandler()

	r := mux.NewRouter()

	r.Use(middleware.RequestLogging)

	r.Handle("/health", healthHandler).Methods("GET")

	address := "localhost:" + cfg.Port

	server := &http.Server{
		Addr:    address,
		Handler: r,
	}

	fmt.Printf("Starting server on %s\n", address)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
