package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tobiasaagaard/localgift-api/config"
	"github.com/tobiasaagaard/localgift-api/middleware"
	"github.com/tobiasaagaard/localgift-api/pkg/response"
)

func main() {
	cfg := config.LoadConfig()

	r := mux.NewRouter()

	r.Use(middleware.RequestLogging)

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		response.Success(w, map[string]string{"Status": "OK"})
	}).Methods("GET")

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
