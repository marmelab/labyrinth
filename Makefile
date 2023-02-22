.PHONY: help
.DEFAULT_GOAL := help

SAVE_ID="_default"

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

run: ## Run the program
	@(go run labyrinth.go -s ${SAVE_ID})

test: ## Run unit tests
	@(go test -race ./...)

clean: ## Cleans all existing saves, use with caution.
	@(rm ${HOME}/.labyrinth/*)
