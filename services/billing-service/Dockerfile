# Use official Golang image as base image
FROM golang:1.19-alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.sum ./

# Download Go dependencies
RUN go mod download

# Copy the entire source code into the container
COPY . .

# Build the Go app
RUN go build -o billing-service .

# Start a new stage from scratch
FROM alpine:latest  

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/billing-service .

# Expose port 3000 to the outside world
EXPOSE 3000

# Run the Go binary
CMD ["./billing-service"]
