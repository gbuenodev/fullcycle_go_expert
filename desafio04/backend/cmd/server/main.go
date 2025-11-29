package main

import (
	"context"
	"log"

	"github.com/gbuenodev/fullcycle_go_expert/desafio04/backend/config"
	"github.com/gbuenodev/fullcycle_go_expert/desafio04/backend/internal/observability"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	ctx := context.Background()
	shutdown, err := observability.InitTelemetry(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize telemetry: %v", err)
	}
	defer func() {
		if err := shutdown(ctx); err != nil {
			log.Printf("Error during telemetry shutdown: %v", err)
		}
	}()

	webServer, err := InitializeWebServer(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize web server: %v", err)
	}

	if err := webServer.Start(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
