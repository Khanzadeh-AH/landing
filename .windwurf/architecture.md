# System Architecture

This repository hosts a full-stack web application with a Go Fiber backend and a SvelteKit frontend. The system follows a simple client-server architecture with clear separation of concerns and environment-driven configuration.

## Overview
- Backend: Go (Fiber) HTTP API providing health, version, and future business endpoints.
- Frontend: SvelteKit SPA/SSR client consuming backend APIs.
- Communication: JSON over HTTP. Public endpoints versioned under `/api`.
- Environments: development, staging, production.

## Components
- Backend (`backend/`)
  - Web server built with Fiber, using middleware for logging, recovery, security headers, CORS, compression, and request IDs.
  - Configuration via environment variables with sensible defaults; supports build metadata (version, commit, build date).
  - Graceful shutdown on SIGINT/SIGTERM.
  - Health (`/healthz`) and version (`/version`) endpoints plus `/api/*` aliases.

- Frontend (`frontend/`)
  - SvelteKit (TypeScript, minimal template) serving the web UI.
  - Development server via Vite; production build emits static assets + server adapter (default).
  - Will consume backend API via a configurable base URL.

## Requests Flow
1. Browser requests frontend pages/assets from SvelteKit dev server (dev) or from deployed static/SSR host (prod).
2. Frontend calls backend API endpoints (e.g., `/api/version`) using fetch.
3. Backend responds with JSON payloads.

## Configuration
- Backend
  - APP_NAME, ENV, PORT
  - VERSION, COMMIT_HASH, BUILD_DATE
- Frontend
  - PUBLIC_API_BASE (planned): e.g., http://localhost:8080 or reverse-proxied path `/api` in prod

## Security & Middleware (Backend)
- recover: panic protection
- requestid: correlation IDs
- logger: structured request logs
- cors: permissive in dev; tighten in prod
- helmet: security headers
- compress: gzip/brotli responses

## Observability (Planned)
- Structured application logs (slog/zerolog)
- Prometheus metrics and health/readiness endpoints
- Request tracing using request IDs

## Deployment (Planned)
- Backend built as a single binary with env-based config.
- Frontend built with SvelteKit and deployed to a static/CDN host or a Node adapter depending on SSR needs.
- Docker images and CI workflows to be added.

## API Versioning Strategy (Planned)
- Public API namespace `/api` with semantic versioning on endpoints as needed (e.g., `/api/v1/*`).
- Backward-compatible changes favored; breaking changes gated behind new versions.
