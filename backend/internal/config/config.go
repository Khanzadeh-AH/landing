package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Config holds application configuration.
type Config struct {
	AppName string
	Env     string
	Port    int

	// Build/Version metadata
	Version    string
	CommitHash string
	BuildDate  string
}

// Load reads configuration from environment variables with defaults.
func Load() Config {
	cfg := Config{
		AppName:    getEnv("APP_NAME", "landing-backend"),
		Env:        strings.ToLower(getEnv("ENV", "development")),
		Port:       getEnvAsInt("PORT", 8080),
		Version:    getEnv("VERSION", ""),
		CommitHash: getEnv("COMMIT_HASH", ""),
		BuildDate:  getEnv("BUILD_DATE", ""),
	}
	return cfg
}

// IsDevelopment returns true if ENV is development.
func (c Config) IsDevelopment() bool {
	return c.Env == "development"
}

// VersionString returns a compact computed version string.
func (c Config) VersionString() string {
	parts := []string{}
	if c.Version != "" {
		parts = append(parts, c.Version)
	}
	if c.CommitHash != "" {
		parts = append(parts, c.CommitHash)
	}
	if c.BuildDate != "" {
		parts = append(parts, c.BuildDate)
	}
	if len(parts) == 0 {
		return "dev"
	}
	return strings.Join(parts, "-")
}

func getEnv(key, def string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return def
}

func getEnvAsInt(key string, def int) int {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			return n
		}
	}
	return def
}

// Addr returns ":port" string
func (c Config) Addr() string { return fmt.Sprintf(":%d", c.Port) }
