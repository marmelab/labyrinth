#!/usr/bin/env bash

mkdir -p logs data/postgres

sudo docker compose \
	--env-file=webapp/.env \
	-f docker-compose.yml \
	-f docker-compose.prod.yml \
	up --pull=always -d
