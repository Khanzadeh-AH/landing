package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib" // register pgx with database/sql
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"

	"landing/backend/ent"
	"landing/backend/internal/config"
)

// OpenClient opens an Ent client using DATABASE_URL from config.
// If cfg.IsDevelopment() is true, it will run auto-migrations.
func OpenClient(ctx context.Context, cfg config.Config) (*ent.Client, error) {
	if cfg.DatabaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is not set")
	}

	sqldb, err := sql.Open("pgx", cfg.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to postgres: %w", err)
	}
	// Validate the connection early to catch DSN or network issues.
	{
		ctxPing, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		if err := sqldb.PingContext(ctxPing); err != nil {
			_ = sqldb.Close()
			return nil, fmt.Errorf("failed to ping database: %w", err)
		}
	}
	drv := entsql.OpenDB(dialect.Postgres, sqldb)
	client := ent.NewClient(ent.Driver(drv))

	// Auto-migrate in development environment
	if cfg.IsDevelopment() {
		if err := client.Schema.Create(ctx); err != nil {
			_ = client.Close()
			return nil, fmt.Errorf("failed running schema migrations: %w", err)
		}
	}

	return client, nil
}
