#!/usr/bin/env bash 

echo "Running docker with user-protected whatever"
docker build -t nats .
docker run -t nats
