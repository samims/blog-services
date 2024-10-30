# Use the official Golang image as a build stage
FROM golang:1.23-alpine AS builder


# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server
RUN CGO_ENABLED=0 GOOS=linux go build -o migrate ./cmd/migrate


# Start a new stage from scratch
FROM alpine:latest
#FROM scratch

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .
COPY --from=builder /app/migrate .
COPY --from=builder /app/migrations ./migrations

## Add wait-for it script
#ADD https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh /usr/local/bin/wait-for-it
#RUN chmod +x /usr/local/bin/wait-for-it

# Create a simple wait-for script
RUN echo '#!/bin/sh\n\
while ! nc -z "$1" "$2"; do\n\
  echo "Waiting for $1:$2...";\n\
  sleep 1;\n\
done\n\
echo "$1:$2 is available"' > /usr/local/bin/wait-for \
    && chmod +x /usr/local/bin/wait-for

# Copy the entrypoint script
COPY entrypoint.sh .
RUN chmod +x entrypoint.sh
# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
#CMD ["./main"]
ENTRYPOINT ["./entrypoint.sh"]