package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tobiasaagaard/localgift-api/config"
)

func main() {
	cfg := config.LoadConfig()

	r := mux.NewRouter()

	address := "localhost:" + cfg.Port
	fmt.Printf("Server running on %s\n", address)
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	err := http.ListenAndServe(address, r)
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
