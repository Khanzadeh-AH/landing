# Application Structure

This document describes the repository layout and the purpose of each directory/file for both backend (Go Fiber) and frontend (SvelteKit).

## Repo layout
```
/                         # monorepo root
  backend/                # Go Fiber backend API
    cmd/
      api/
        main.go           # app entrypoint
    internal/
      config/
        config.go         # env config & version metadata
      handlers/
        health.go         # health/version handlers
      middleware/
        middleware.go     # recover, requestid, logger, cors, helmet, compress
      routes/
        routes.go         # route registration
    .env.example          # sample env for local dev
    .gitignore
    go.mod
    README.md

  frontend/               # SvelteKit frontend (TypeScript)
    src/
      app.d.ts
      app.css
      routes/
        +layout.svelte
        +page.svelte      # home page (template)
    static/               # static assets
    package.json
    svelte.config.js
    vite.config.ts
    tsconfig.json
    .npmrc (optional)

  .windwurf/              # project docs for Windsurf usage
    architecture.md       # system architecture overview
    app-structure.md      # this document
```

## Backend
- Language: Go 1.24.6, Fiber v2.
- Entrypoint: `backend/cmd/api/main.go`.
- Endpoints: `/healthz`, `/version`, and `/api/*` aliases.
- Config via env: `APP_NAME`, `ENV`, `PORT`, `VERSION`, `COMMIT_HASH`, `BUILD_DATE`.
- Run: `cd backend && go run ./cmd/api`.
- Build: `cd backend && go build -o bin/api ./cmd/api`.

## Frontend
- Framework: SvelteKit (TypeScript, minimal template).
- Dev: `cd frontend && npm run dev` (defaults to http://localhost:5173).
- Build: `cd frontend && npm run build`.
- Preview: `cd frontend && npm run preview` (defaults to http://localhost:4173).
- Planned env: `PUBLIC_API_BASE` for API base URL.

## Conventions
- API routes under `/api` namespace; future versioning as `/api/v1/*`.
- Use request IDs for tracing across services.
- Prefer environment-based configuration; no secrets in VCS.

## Roadmap (high level)
- Add Makefile and Dockerfiles for both apps.
- Add CI pipelines (lint, test, build, scan).
- Add dev proxy from frontend to backend.
- Add sample API consumption page in frontend.
- Add OpenAPI spec & typed client generation (planned).
