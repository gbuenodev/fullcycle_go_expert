package main

import (
	"log"

	"github.com/gbuenodev/fullcycle_go_expert/desafio04/backend/config"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	webServer, err := InitializeWebServer(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize web server: %v", err)
	}

	if err := webServer.Start(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
