package db

import (
    "context"
    "log"
    "strings"
    
    "landing/backend/ent"
    "landing/backend/internal/config"
)

// EnsureClient returns a live Ent client. If the global client is nil, it will
// initialize a new one and set it globally. If the existing client appears closed,
// it will attempt to reopen using a background context.
func EnsureClient(ctx context.Context, cfg config.Config) (*ent.Client, error) {
    if cli := GlobalClient(); cli != nil {
        // Probe with a background context; if it fails with a closed DB error, reopen.
        if _, err := cli.Blog.Query().Limit(1).Count(context.Background()); err != nil {
            // Only reopen on indicative errors to avoid reopening on transient query errors.
            msg := err.Error()
            if strings.Contains(msg, "database is closed") || strings.Contains(msg, "driver: bad connection") {
                log.Printf("db.EnsureClient: existing client seems closed, reopening: %v", err)
            } else {
                return cli, nil
            }
        } else {
            return cli, nil
        }
    }
    // Use a background context for initialization to avoid request-scoped cancellations.
    newCli, err := OpenClient(context.Background(), cfg)
    if err != nil {
        return nil, err
    }
    SetGlobalClient(newCli)
    log.Printf("db.EnsureClient: initialized Ent client")
    return newCli, nil
}
