package db

import (
	"context"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib" // register pgx with database/sql

	"landing/backend/ent"
	"landing/backend/internal/config"
)

// OpenClient opens an Ent client using DATABASE_URL from config.
// If cfg.IsDevelopment() is true, it will run auto-migrations.
func OpenClient(ctx context.Context, cfg config.Config) (*ent.Client, error) {
	if cfg.DatabaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is not set")
	}

	client, err := ent.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to postgres: %w", err)
	}

	// Auto-migrate in development environment
	if cfg.IsDevelopment() {
		if err := client.Schema.Create(ctx); err != nil {
			_ = client.Close()
			return nil, fmt.Errorf("failed running schema migrations: %w", err)
		}
	}

	return client, nil
}
