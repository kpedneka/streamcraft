# syntax=docker/dockerfile:1

# --- Stage 1: Build the backend binary ---
FROM golang:1.27-alpine AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to cache dependencies
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download

# Copy the source code into the container
COPY . .

# Build the Go app as a statically-linked binary
# CGO_ENABLED=0 removes dependency on dynamic C libraries
# GOOS=linux ensures the binary runs natively inside Linux containers
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o streamcraft ./cmd/api

# --- Stage 2: Final lightweight image ---
FROM alpine:3.20 AS runtime

# Install ca-certificates in case your app makes HTTPS requests
RUN apk --no-cache add ca-certificates

# Set the current working directory inside the container
WORKDIR /app
COPY --from=builder /app/streamcraft     /app/streamcraft

# Expose the application port
EXPOSE 8080

# Run the compiled binary
CMD ["./streamcraft"]
