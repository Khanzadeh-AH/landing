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

	// Database
	DatabaseURL string

	// API key for protecting endpoints
	APIKey string

	// Build/Version metadata
	Version    string
	CommitHash string
	BuildDate  string

	// Site metadata for blog placeholder replacements
	SiteName             string
	SiteBaseURL          string
	SiteLogo             string
	DefaultFeaturedImage string
	AuthorName           string
	AuthorBio            string
}

// Load reads configuration from environment variables with defaults.
func Load() Config {
	cfg := Config{
		AppName:     getEnv("APP_NAME", "landing-backend"),
		Env:         strings.ToLower(getEnv("ENV", "development")),
		Port:        getEnvAsInt("PORT", 8080),
		DatabaseURL: getEnv("DATABASE_URL", ""),
		APIKey:      getEnv("API_KEY", ""),
		Version:     getEnv("VERSION", ""),
		CommitHash:  getEnv("COMMIT_HASH", ""),
		BuildDate:   getEnv("BUILD_DATE", ""),

		// Site metadata
		SiteName:             getEnv("SITE_NAME", "Landing"),
		SiteBaseURL:          getEnv("SITE_BASE_URL", ""),
		SiteLogo:             getEnv("SITE_LOGO", "/favicon.ico"),
		DefaultFeaturedImage: getEnv("FEATURED_IMAGE", "/og-default.jpg"),
		AuthorName:           getEnv("AUTHOR_NAME", "Admin"),
		AuthorBio:            getEnv("AUTHOR_BIO", ""),
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
