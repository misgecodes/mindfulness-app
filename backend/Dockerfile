# ---------- Build stage ----------
FROM golang:tip-alpine AS builder

WORKDIR /app

# Install git (needed for go mod download sometimes)
RUN apk add --no-cache git

# Copy go mod files first (better caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

# ---------- Runtime stage ----------
FROM alpine:latest

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/app .

# Expose port
EXPOSE 8080

# Run app
CMD ["./app"]
