# ---- Build Stage ----
FROM golang:1.24.1-alpine AS builder

# Set working directory inside the container
WORKDIR /app

# Copy Go modules & dependencies first (for caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Change directory to cmd/app and build the Go binary
RUN cd cmd/app && go build -o /app/main .

# ---- Final Stage ----
FROM alpine:latest

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Set environment variable inside the container
ENV APP_ENV=development

# Expose the application port (change if needed)
EXPOSE 8080

# Run the application
CMD ["./main"]