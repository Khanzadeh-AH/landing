# Backend (Go Fiber)

A production-ready scaffold using Go Fiber with health and version endpoints, structured layout, and sensible middleware defaults.

## Requirements
- Go 1.24.6 (toolchain). Install from https://go.dev/dl/

## Structure
```
backend/
  cmd/
    api/
      main.go
  internal/
    config/
      config.go
    handlers/
      health.go
    middleware/
      middleware.go
    routes/
      routes.go
  .gitignore
  README.md
  .env (optional)
```

## Configuration (env)
- PORT: Server port (default: 8080)
- ENV: runtime environment: development|staging|production (default: development)
- APP_NAME: Application name (default: landing-backend)
- VERSION: Semantic version e.g. 0.1.0 (optional)
- COMMIT_HASH: Git commit SHA (optional)
- BUILD_DATE: ISO-8601 build date (optional)

Place variables in `.env` for local development (loaded automatically if present).

## Run locally
Initialize module and dependencies (choose your preferred module path):

```bash
# from repo root
cd backend

# pick a module path; you can change this later
MODULE_PATH="landing/backend"
go mod init "$MODULE_PATH"

# Fetch latest Fiber and middleware
go get github.com/gofiber/fiber/v2 \
       github.com/gofiber/fiber/v2/middleware/logger \
       github.com/gofiber/fiber/v2/middleware/recover \
       github.com/gofiber/fiber/v2/middleware/cors \
       github.com/gofiber/helmet/v2 \
       github.com/gofiber/compress/v2 \
       github.com/gofiber/requestid/v3 \
       github.com/joho/godotenv

go mod tidy

# run
go run ./cmd/api
```

Then open: http://localhost:8080/healthz and http://localhost:8080/version

## Build
```bash
PORT=8080 ENV=production go build -o bin/api ./cmd/api
```

## Notes
- Uses graceful shutdown (SIGINT/SIGTERM).
- Default middlewares: recover, requestid, logger, cors, helmet, compress.
- Health and version endpoints return JSON.
