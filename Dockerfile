# syntax=docker/dockerfile:1

# Use the official Golang image as the base image for the build stage
FROM golang:1.22-alpine AS builder

# Set the Current Working Directory inside the builder container
WORKDIR /app

# Copy go.mod and go.sum files to the workspace
COPY app/go.mod app/go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code to the workspace
COPY app/ .

# Build the Go application
RUN go build -o /hka-server-login ./cmd/main

# Use a minimal base image to reduce the final image size
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the pre-built binary from the builder stage
COPY --from=builder /hka-server-login .
COPY app/config/config.json .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./hka-server-login", "--configPath", "./config.json"]
