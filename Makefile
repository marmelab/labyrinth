.PHONY: help
.DEFAULT_GOAL := help

SAVE_ID="_default"
BOARD_SIZE="7"
PLAYER_COUNT="1"

# This is the hash used to tag Docker images
COMMIT_HASH=$(shell git rev-parse HEAD)
DOCKER_IMAGE_NAMESPACE=jonathanmarmelab
DOCKER_IMAGE_DOMAIN_API=${DOCKER_IMAGE_NAMESPACE}/labyrinth-domain-api
DOCKER_IMAGE_WEBAPP=${DOCKER_IMAGE_NAMESPACE}/labyrinth-webapp
DOCKER_IMAGE_PROXY=${DOCKER_IMAGE_NAMESPACE}/labyrinth-proxy

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

setup-env:										## Setup the environment
	@(${MAKE} -C webapp setup-env)

install:										## Install dependencies
	@(${MAKE} -C domain install)
	@(${MAKE} -C webapp install)

run: develop									## Run the program for development, alias of develop

develop: 										## Run the program for development
	@(mkdir -p logs)
	docker compose \
		-f docker-compose.yml \
		-f docker-compose.dev.yml \
		up --build

develop-config: 								## Dumps the docker development compose file with environment set
	docker compose \
		-f docker-compose.yml \
		-f docker-compose.dev.yml \
		config

production-image:
	docker build \
		-f domain/api/Dockerfile \
		-t ${DOCKER_IMAGE_DOMAIN_API}:${COMMIT_HASH} \
		-t ${DOCKER_IMAGE_DOMAIN_API}:latest \
		.

	docker build \
		-f webapp/Dockerfile \
		-t ${DOCKER_IMAGE_WEBAPP}:${COMMIT_HASH} \
		-t ${DOCKER_IMAGE_WEBAPP}:latest \
		.

	docker build \
		-f proxy/Dockerfile \
		-t ${DOCKER_IMAGE_PROXY}:${COMMIT_HASH} \
		-t ${DOCKER_IMAGE_PROXY}:latest \
		.

production-image-push: production-image			## Push production images to Docker Hub
	docker image push --all-tags ${DOCKER_IMAGE_DOMAIN_API}
	docker image push --all-tags ${DOCKER_IMAGE_WEBAPP}
	docker image push --all-tags ${DOCKER_IMAGE_PROXY}

production-deploy: production-image-push		## Deploy production to AWS
	scp \
		-i .secrets/labyrinth-ed25519.pem \
		docker-compose.yml docker-compose.prod.yml scripts/run-production.sh \
		ubuntu@ec2-13-37-240-163.eu-west-3.compute.amazonaws.com:~

	ssh \
	 	-i .secrets/labyrinth-ed25519.pem \
	 	ubuntu@ec2-13-37-240-163.eu-west-3.compute.amazonaws.com \
	 	'echo "TAG=${COMMIT_HASH}" > .env'
	
	ssh \
	 	-i .secrets/labyrinth-ed25519.pem \
	 	ubuntu@ec2-13-37-240-163.eu-west-3.compute.amazonaws.com \
	 	'./run-production.sh'

production: production-image					## Run the program for production
	@(mkdir -p logs)
	TAG=${COMMIT_HASH} docker compose \
		--pull always \
		-f docker-compose.yml \
		-f docker-compose.prod.yml \
		up

test: 											## Run unit tests
	@(${MAKE} -C domain test)
	@(${MAKE} -C webapp test)

cli-run: 										## Run the CLI version of the labyrinth.
	@(${MAKE} -C domain run)

cli-clean: 										## Cleans all existing saves, use with caution.
	@(rm ${HOME}/.labyrinth/*)
