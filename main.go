package main

import (
	"fmt"

	"github.com/tobiasaagaard/localgift-api/config"
)

func main() {
	cfg := config.LoadConfig()
	fmt.Printf("Starting LocalGift API on port %s with database %s\n", cfg.Port, cfg.DatabaseURL)
}
