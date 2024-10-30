#!/bin/sh

echo "Running blog_db migrations..."

./migrate -migration-dir=./migrations

echo "Starting the application...."

exec ./main