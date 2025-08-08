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

	"landing/backend/internal/config"
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
