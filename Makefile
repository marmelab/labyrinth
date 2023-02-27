.PHONY: help
.DEFAULT_GOAL := help

SAVE_ID="_default"
BOARD_SIZE="7"
PLAYER_COUNT="1"

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

run: 									## Run the program
	@(echo "TODO")

test: test-domain						## Run unit tests

test-domain:							## Run unit tests for the domain directory
	@(${MAKE} -C domain test)

cli-run: 								## Run the CLI version of the labyrinth.
	@(${MAKE} -C domain run)

cli-clean: 								## Cleans all existing saves, use with caution.
	@(rm ${HOME}/.labyrinth/*)
