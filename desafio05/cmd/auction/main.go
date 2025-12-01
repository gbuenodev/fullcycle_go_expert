package main

import (
	"context"
	"log"

	"github.com/gbuenodev/fullcycle_go_expert/desafio05/configs/db/mongodb"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	if err := godotenv.Load("cmd/auction/.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
		return
	}

	db, err := mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer db.Client().Disconnect(ctx)
}
