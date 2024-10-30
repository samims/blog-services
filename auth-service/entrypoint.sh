#!/bin/sh



# RUN migrations
echo "Running database migrations..."
./migrate -migration-dir=./migrations

# Start the main app
echo "Starting the application..."

exec ./main
