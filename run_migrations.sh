#!/bin/bash

NETWORK_NAME=$(docker network ls | grep hacknet | awk '{print $2}')

if [ -f .env ]; then
  export $(grep -v '^#' .env | sed 's/^/export /')
fi

docker run --rm \
  -v $(pwd)/migrations:/migrations \
  --network $NETWORK_NAME \
  migrate/migrate \
  -path /migrations -database "postgres://$DB_USER:$DB_PASSWORD@db:5432/$DB_NAME?sslmode=disable" up
