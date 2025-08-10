package db

import (
	"context"
	"landing/backend/ent"
	"landing/backend/internal/config"

	"github.com/gofiber/fiber/v2"
)

// ClientFromCtx retrieves the Ent client injected into the request context.
func ClientFromCtx(c *fiber.Ctx) *ent.Client {
	// Prefer request-scoped client if present and typed correctly.
	if v := c.Locals("ent"); v != nil {
		if client, ok := v.(*ent.Client); ok {
			return client
		}
	}

	// Ensure a live global client (auto-reopen if the underlying connection was closed),
	// then attach it to the request context for downstream handlers.
	cfg := config.Load()
	cli, err := EnsureClient(context.Background(), cfg)
	if err != nil {
		// As a last resort, return whatever global client exists (may be nil)
		return GlobalClient()
	}
	c.Locals("ent", cli)
	return cli
}
