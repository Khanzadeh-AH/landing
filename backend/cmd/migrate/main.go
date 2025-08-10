package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"

	"landing/backend/internal/config"
	"landing/backend/internal/db"
)

func main() {
	// Load .env if present
	_ = godotenv.Load()

	cfg := config.Load()
	ctx := context.Background()

	client, err := db.OpenClient(ctx, cfg)
	if err != nil {
		log.Fatalf("migrate: database initialization failed: %v", err)
	}
	// Ensure DB is closed on exit
	db.EnableDBClose()
	defer func() {
		if err := client.Close(); err != nil {
			log.Printf("migrate: error closing db client: %v", err)
		}
	}()

	// Run schema migrations (idempotent)
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("migrate: schema migration failed: %v", err)
	}

	log.Println("migrate: schema migration completed successfully")
}
