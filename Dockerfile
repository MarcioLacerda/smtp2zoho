# Stage 1: Build
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy project files
COPY . .

# Install git in case dependencies require it
RUN apk add --no-cache git

# Build for current architecture (your server is ARM64)
RUN go build -o smtp2zoho

# Stage 2: Minimal runtime image
FROM alpine:3.19

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/smtp2zoho /usr/local/bin/smtp2zoho

# Expose SMTP port
EXPOSE 1025

# Run binary
CMD ["smtp2zoho"]
