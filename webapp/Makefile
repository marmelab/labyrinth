.PHONY: help
.DEFAULT_GOAL := help

E2E_NO_HEADLESS="0"
E2E_DEVTOOLS="0"

JWT_DIR=config/jwt
JWT_PUB_PATH="${JWT_DIR}/public.pem"
JWK_KEY_PATH="${JWT_DIR}/jwk.pub"

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

setup-env:					## Setup the environment
	composer install
	cp -n .env.dist .env
	php bin/console lexik:jwt:generate-keypair --skip-if-exists
	cat ${JWT_PUB_PATH} | docker run -i danedmunds/pem-to-jwk:latest > ${JWK_KEY_PATH}
	sed -i -e 's/}/,"alg":"RS256","use":"sig"}/g' ${JWK_KEY_PATH}

install: setup-env			## Install dependencies

create-admin:
	php bin/console security:hash-password
	php bin/console run:psql

test: install				## Run PHP tests
	SYMFONY_DEPRECATIONS_HELPER=weak php bin/phpunit tests/Unit

test-e2e:					## Run end to end tests
	@(mkdir -p ../logs/e2e)
	SYMFONY_DEPRECATIONS_HELPER=weak \
	PANTHER_NO_HEADLESS=${E2E_NO_HEADLESS} \
	PANTHER_ERROR_SCREENSHOT_DIR=../logs/e2e \
	PANTHER_DEVTOOLS=${E2E_DEVTOOLS} \
		php bin/phpunit tests/E2E
