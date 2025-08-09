package db

import (
    "context"
    "log"
    "time"

    "landing/backend/ent"
    "landing/backend/internal/config"
)

// EnsureClient returns a live Ent client. If the global client is nil or appears closed,
// it will attempt to reopen using the provided config and set it globally.
func EnsureClient(ctx context.Context, cfg config.Config) (*ent.Client, error) {
    if cli := GlobalClient(); cli != nil {
        // Probe the connection with a cheap query to detect a closed DB pool.
        probeCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
        defer cancel()
        if _, probeErr := cli.Blog.Query().Limit(1).Count(probeCtx); probeErr == nil {
            return cli, nil
        } else {
            // If database is closed or probe failed, try to reopen.
            log.Printf("db.EnsureClient: existing client seems unusable, reopening: %v", probeErr)
        }
    }
    newCli, err := OpenClient(ctx, cfg)
    if err != nil {
        return nil, err
    }
    SetGlobalClient(newCli)
    log.Printf("db.EnsureClient: opened new Ent client")
    return newCli, nil
}
