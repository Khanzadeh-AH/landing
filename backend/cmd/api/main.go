// @title Landing Backend API
// @version 1.0
// @description OpenAPI documentation for Landing backend.
// @BasePath /api
//
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-API-Key
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
	_ "landing/backend/docs"

	"landing/backend/internal/config"
	"landing/backend/internal/db"
	"landing/backend/internal/middleware"
	"landing/backend/internal/routes"
)

func main() {
	// Load .env if present (ignore error if not present)
	_ = godotenv.Load()

	cfg := config.Load()

	app := fiber.New(fiber.Config{
		AppName:               cfg.AppName,
		EnablePrintRoutes:     cfg.IsDevelopment(),
		DisableStartupMessage: false,
	})

	// Initialize database (Ent)
	{
		ctx := context.Background()
		client, err := db.OpenClient(ctx, cfg)
		if err != nil {
			log.Fatalf("database initialization failed: %v", err)
		}
        // Set global client fallback
        db.SetGlobalClient(client)
		// Seed development data
		if cfg.IsDevelopment() {
			db.SeedDev(ctx, client)
		}
        // Attach the initialized Ent client to each request (kept open for app lifetime).
        app.Use(func(c *fiber.Ctx) error {
            c.Locals("ent", client)
            return c.Next()
        })
		// Ensure DB is closed on app shutdown
		app.Hooks().OnShutdown(func() error {
			log.Println("OnShutdown: closing Ent DB client...")
			// Allow the wrapped driver to actually close at shutdown time.
			db.EnableDBClose()
			if err := client.Close(); err != nil {
				log.Printf("error closing db client: %v", err)
				return err
			}
			log.Println("OnShutdown: Ent DB client closed")
			return nil
		})
	}

	// Global middleware
	middleware.Register(app, cfg)

	// Routes
	routes.Register(app, cfg)

	// Start server with graceful shutdown
	addr := fmt.Sprintf(":%d", cfg.Port)

	srvErr := make(chan error, 1)
	go func() {
		log.Printf("starting %s on %s (env=%s)\n", cfg.AppName, addr, cfg.Env)
		srvErr <- app.Listen(addr)
	}()

	// Listen for SIGINT/SIGTERM
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-quit:
		log.Printf("received signal: %s, shutting down...\n", sig)
	case err := <-srvErr:
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Printf("graceful shutdown failed: %v", err)
	}

	log.Println("server stopped")
}
