#!/usr/bin/env bash

source .env

mkdir -p logs data

sudo docker compose \
	-f docker-compose.yml \
	-f docker-compose.prod.yml \
	up --pull always -d
