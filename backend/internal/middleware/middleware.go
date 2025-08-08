package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/helmet/v2"

	"landing/backend/internal/config"
)

// Register attaches global middleware to the Fiber app.
func Register(app *fiber.App, cfg config.Config) {
	// Panic recovery
	app.Use(recover.New())

	// Unique request ID for tracing
	app.Use(requestid.New())

	// Structured logging
	app.Use(logger.New(logger.Config{
		Format:     "${time} ${ip} ${status} ${method} ${path} ${latency} req_id=${locals:requestid}\n",
		TimeZone:   "Local",
		TimeFormat: time.RFC3339,
	}))

	// CORS (permissive defaults for dev; tighten in prod)
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,PATCH,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		ExposeHeaders:    "",
		AllowCredentials: false,
	}))

	// Security headers
	app.Use(helmet.New())

	// Response compression
	app.Use(compress.New())
}
