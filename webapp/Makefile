.PHONY: help
.DEFAULT_GOAL := help

E2E_NO_HEADLESS="0"
E2E_DEVTOOLS="0"

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

setup-env:					## Setup the environment
	composer install
	cp -n .env.dist .env
	php bin/console lexik:jwt:generate-keypair

install: setup-env			## Install dependencies
	

test: install				## Run PHP tests
	SYMFONY_DEPRECATIONS_HELPER=weak php bin/phpunit tests/Unit

test-e2e:					## Run end to end tests
	@(mkdir -p ../logs/e2e)
	SYMFONY_DEPRECATIONS_HELPER=weak \
	PANTHER_NO_HEADLESS=${E2E_NO_HEADLESS} \
	PANTHER_ERROR_SCREENSHOT_DIR=../logs/e2e \
	PANTHER_DEVTOOLS=${E2E_DEVTOOLS} \
		php bin/phpunit tests/E2E
