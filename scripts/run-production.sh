#!/usr/bin/env bash

source .env

mkdir -p logs

sudo docker compose \
	-f docker-compose.yml \
	-f docker-compose.prod.yml \
	up --pull always -d
