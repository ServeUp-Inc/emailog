#!/bin/bash

# Run MySQL instance
# This command requires podman machine to be running (`podman machine start`)
# Image tag is oracle because it supports arm64/v8
podman run -d \
  --name sqldb \
  -p 3306:3306 \
  -e MYSQL_ROOT_PASSWORD=root \
  -e MYSQL_USER=testuser \
  -e MYSQL_PASSWORD=testpass \
  -e MYSQL_DATABASE=testdb \
  mysql:oracle
