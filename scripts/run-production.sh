#!/usr/bin/env bash

mkdir -p logs data/postgres certbot/{conf,www}

sudo docker compose \
	--env-file=webapp/.env \
	-f docker-compose.yml \
	-f docker-compose.prod.yml \
	up --pull=always -d
