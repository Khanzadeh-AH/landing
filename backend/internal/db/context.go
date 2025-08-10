package db

import (
    "landing/backend/ent"

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
    // Fallback to global client without trying to reopen.
    return GlobalClient()
}
