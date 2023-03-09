.PHONY: help
.DEFAULT_GOAL := help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

install:					## Install dependencies
	@(if [ -f ${HOME}/.nvm/nvm.sh ] ; \
	then \
		. ${HOME}/.nvm/nvm.sh && nvm install && nvm use; \
	fi)

	@npm install

test: install				## Run Vitest
	@npm test
