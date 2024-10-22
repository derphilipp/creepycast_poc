# Stage 1: Build the Go application
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod ./

# Download the Go modules
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o podcast-server .

# Stage 2: Create the final image
FROM alpine:latest

# Install ffmpeg - tzdata required for local timezone usage
RUN apk add --no-cache ffmpeg tzdata

# Set the working directory inside the container
WORKDIR /app

# Copy the built Go application from the builder stage
COPY --from=builder /app/podcast-server .

# Copy the variations directory containing the MP3 files
COPY ./variations ./variations

# Expose the port the application runs on
EXPOSE 8080

# Command to run the application
CMD ["./podcast-server"]
