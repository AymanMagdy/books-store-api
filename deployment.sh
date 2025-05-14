#!/bin/bash

set -e

docker build -t books-app .

docker compose -d

sleep 5 # wait 5 seconds till the service is up and running

echo "#### Sending test request to http://localhost:8080 ####"

response=$(curl -s -o /dev/null -w "%{http_code}" http://localhost:8080)

if [ "$response" -eq 200 ]; then
  echo "API is up and running!"
else
  echo "HTTP Failed status: $response"
fi