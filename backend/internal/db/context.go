package db

import (
	"landing/backend/ent"

	"github.com/gofiber/fiber/v2"
)

// ClientFromCtx retrieves the Ent client injected into the request context.
func ClientFromCtx(c *fiber.Ctx) *ent.Client {
	v := c.Locals("ent")
	if v == nil {
		return nil
	}
	if client, ok := v.(*ent.Client); ok {
		return client
	}
	return nil
}
