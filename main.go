package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tobiasaagaard/localgift-api/config"
)

func main() {
	cfg := config.LoadConfig()

	r := mux.NewRouter()

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

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
