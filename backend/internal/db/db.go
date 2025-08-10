package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq" // register lib/pq with database/sql
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

	// Use lib/pq driver
	sqldb, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to postgres: %w", err)
	}
	// Reasonable pool defaults; adjust as needed via env in the future.
	sqldb.SetMaxOpenConns(10)
	sqldb.SetMaxIdleConns(5)
	sqldb.SetConnMaxLifetime(1 * time.Hour)
	// Validate the connection early to catch DSN or network issues.
	{
		ctxPing, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		if err := sqldb.PingContext(ctxPing); err != nil {
			_ = sqldb.Close()
			return nil, fmt.Errorf("failed to ping database: %w", err)
		}
	}
	base := entsql.OpenDB(dialect.Postgres, sqldb)
    // Prevent accidental closure during runtime; allow closing only on shutdown.
    wrapped := wrapKeepOpen(base)
	client := ent.NewClient(ent.Driver(wrapped))

	// Auto-migrate in development environment
	if cfg.IsDevelopment() {
		if err := client.Schema.Create(ctx); err != nil {
			_ = client.Close()
			return nil, fmt.Errorf("failed running schema migrations: %w", err)
		}
	}

	return client, nil
}
