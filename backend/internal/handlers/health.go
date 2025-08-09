package handlers

import (
	"github.com/gofiber/fiber/v2"

	"landing/backend/internal/config"
)

// HealthHandler returns a simple OK for liveness/readiness checks.
// @Summary Health check
// @Tags system
// @Produce json
// @Success 200 {object} map[string]string
// @Router /healthz [get]
func HealthHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "ok",
	})
}

// VersionHandler returns a handler that reports version/build metadata.
func VersionHandler(cfg config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"name":        cfg.AppName,
			"env":         cfg.Env,
			"version":     cfg.Version,
			"commit_hash": cfg.CommitHash,
			"build_date":  cfg.BuildDate,
			"composed":    cfg.VersionString(),
		})
	}
}
