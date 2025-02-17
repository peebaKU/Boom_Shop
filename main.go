package main

import (
	"context"
	"github/peebaKU/Boom_Shop/config"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	_ = ctx
	// Initialize config
	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal("Error: .env file path is required")
		}
		return os.Args[1]
	}())
	log.Println(cfg)
}
