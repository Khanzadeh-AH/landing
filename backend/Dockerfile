# syntax=docker/dockerfile:1

# -----------------------------
# Builder image
# -----------------------------
ARG GO_VERSION=1.24.6
FROM golang:${GO_VERSION}-bookworm AS builder

WORKDIR /src

# Enable static build (no CGO)
ENV CGO_ENABLED=0

# Pre-cache dependencies
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# Copy the rest of the source
COPY . .

# Build the API binary
# You can pass custom values on build with:
#   docker build --build-arg VERSION=$(git describe --tags --always) \
#                --build-arg COMMIT_HASH=$(git rev-parse --short HEAD) \
#                --build-arg BUILD_DATE=$(date -u +"%Y-%m-%dT%H:%M:%SZ") \
#                -t landing-backend:latest .
ARG APP_NAME=landing-backend
ARG VERSION=0.0.0
ARG COMMIT_HASH=unknown
ARG BUILD_DATE=unknown

# Produce a small, reproducible binary
RUN --mount=type=cache,target=/root/.cache/go-build \
    go build -trimpath -ldflags "-s -w" -o /out/server ./cmd/api && \
    go build -trimpath -ldflags "-s -w" -o /out/migrate ./cmd/migrate

# -----------------------------
# Runtime image
# -----------------------------
# Use distroless for minimal, secure runtime with CA certs
FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app

# Carry version/build metadata into the runtime image as environment variables
ARG APP_NAME=landing-backend
ARG VERSION=0.0.0
ARG COMMIT_HASH=unknown
ARG BUILD_DATE=unknown
ENV APP_NAME=${APP_NAME} \
    VERSION=${VERSION} \
    COMMIT_HASH=${COMMIT_HASH} \
    BUILD_DATE=${BUILD_DATE} \
    ENV=production \
    PORT=8080

# Copy binary from builder
COPY --from=builder /out/server /app/server
COPY --from=builder /out/migrate /app/migrate

# Run as non-root user provided by the distroless image
USER nonroot:nonroot

# Expose application port (kept in sync with default PORT=8080)
EXPOSE 8080

# Start the server
ENTRYPOINT ["/app/server"]
