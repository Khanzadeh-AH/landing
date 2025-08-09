package db

import "landing/backend/ent"

var globalClient *ent.Client

// SetGlobalClient sets the process-wide Ent client.
func SetGlobalClient(c *ent.Client) { globalClient = c }

// GlobalClient returns the process-wide Ent client if set.
func GlobalClient() *ent.Client { return globalClient }
