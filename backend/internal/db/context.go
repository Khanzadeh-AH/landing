package db

import (
	"landing/backend/ent"

	"github.com/gofiber/fiber/v2"
)

// ClientFromCtx retrieves the Ent client injected into the request context.
func ClientFromCtx(c *fiber.Ctx) *ent.Client {
	v := c.Locals("ent")
	if v == nil {
		// Fallback to global client if request-scoped client is missing
		return GlobalClient()
	}
	if client, ok := v.(*ent.Client); ok {
		return client
	}
	// Fallback to global client on type mismatch
	return GlobalClient()
}
