.PHONY: help
.DEFAULT_GOAL := help

SAVE_ID="_default"
BOARD_SIZE="7"
PLAYER_COUNT="1"

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

setup-env:								## Setup the environment
	@(${MAKE} -C webapp setup-env)

install:								## Install dependencies
	@(${MAKE} -C domain install)
	@(${MAKE} -C webapp install)

run: develop							## Run the program for development, alias of develop

develop: 								## Run the program for development
	@(mkdir -p logs)
	docker compose \
		-f docker-compose.yml \
		-f docker-compose.dev.yml \
		up --build

develop-config: 						## Dumps the docker development compose file with environment set
	docker compose \
		-f docker-compose.yml \
		-f docker-compose.dev.yml \
		config

production: 							## Run the program for production
	@(mkdir -p logs)
	docker compose \
		-f docker-compose.yml \
		-f docker-compose.prod.yml \
		up --build

production-config: 						## Dumps the docker production compose file with environment set
	docker compose \
		-f docker-compose.yml \
		-f docker-compose.prod.yml \
		config

test: 									## Run unit tests
	@(${MAKE} -C domain test)
	@(${MAKE} -C webapp test)

cli-run: 								## Run the CLI version of the labyrinth.
	@(${MAKE} -C domain run)

cli-clean: 								## Cleans all existing saves, use with caution.
	@(rm ${HOME}/.labyrinth/*)
