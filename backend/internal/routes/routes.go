package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	"landing/backend/internal/config"
	"landing/backend/internal/handlers"
	"landing/backend/internal/middleware"
)

// Register wires all HTTP routes.
func Register(app *fiber.App, cfg config.Config) {
	// Swagger UI at /swagger/index.html
	app.Get("/swagger/*", swagger.HandlerDefault)

	api := app.Group("/api")

	// health
	api.Get("/healthz", handlers.HealthHandler)

	// version & root aliases
	api.Get("/version", handlers.VersionHandler(cfg))

	// Protect subsequent /api routes with API key (if configured)
	api.Use(middleware.APIKey(cfg))

	// blogs
	api.Post("/blogs", handlers.CreateBlogHandler)
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
