package routes

import (
	"github.com/gofiber/fiber/v2"

	"landing/backend/internal/config"
	"landing/backend/internal/handlers"
)

// Register wires all HTTP routes.
func Register(app *fiber.App, cfg config.Config) {
	api := app.Group("/api")

	// health
	api.Get("/healthz", handlers.HealthHandler)

	// version & root aliases
	api.Get("/version", handlers.VersionHandler(cfg))

	// blogs
	api.Get("/blogs", handlers.ListBlogsHandler)
	api.Get("/blogs/:path", handlers.GetBlogByPathHandler)

	// convenience root routes
	app.Get("/healthz", handlers.HealthHandler)
	app.Get("/version", handlers.VersionHandler(cfg))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"name":    cfg.AppName,
			"env":     cfg.Env,
			"version": cfg.VersionString(),
		})
	})
}
