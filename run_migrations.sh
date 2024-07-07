#!/bin/bash

# Determine the Docker network name
NETWORK_NAME=$(docker network ls | grep hacknet | awk '{print $2}')

# Check if .env file exists and load environment variables
if [ ! -f ".env" ]; then
    echo ".env file not found."
    exit 1
else
    export $(grep -v '^#' .env | sed 's/^/export /')
fi

# Validate environment variables
if [ -z "$DB_USER" ] || [ -z "$DB_PASSWORD" ] || [ -z "$DB_NAME" ]; then
    echo "Database credentials are missing."
    exit 1
fi

# Run the migration container
docker run --rm \
  -v "$(pwd)/migrations:/migrations" \
  --network hackaton_hacknet \
  migrate/migrate \
  -path /migrations -database "postgres://$DB_USER:$DB_PASSWORD@db:5432/$DB_NAME?sslmode=disable" up
