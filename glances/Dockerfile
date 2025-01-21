# Use the official Golang image to create a build artifact.
FROM golang:1.22 AS builder

# Copy local code to the container image.
WORKDIR /app
COPY . .

# Build the command inside the container with static linking.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

# Use a minimal Debian image and install ca-certificates.
FROM debian:bullseye-slim
WORKDIR /root/

# Install ca-certificates to enable HTTPS verification.
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/server .

# Run the web service on container startup.
CMD ["./server"]
