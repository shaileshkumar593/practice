# ============================================================
# Stage 1: Build
# ============================================================

FROM golang:1.25-alpine AS builder

WORKDIR /app

# Install certificates and git
RUN apk add --no-cache ca-certificates git

# Copy dependency files first
# This allows Docker to cache dependencies
COPY go.mod go.sum ./

# Download Go dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" \
    -o /app/bin/api \
    ./cmd/api


# ============================================================
# Stage 2: Runtime
# ============================================================

FROM alpine:3.20

WORKDIR /app

# Install CA certificates
RUN apk add --no-cache ca-certificates

# Copy compiled binary from builder
COPY --from=builder /app/bin/api /app/api

# Application port
EXPOSE 8080

# Run application
ENTRYPOINT ["/app/api"]

