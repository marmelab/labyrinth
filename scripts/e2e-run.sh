#!/usr/bin/env bash

if [ -f cypress.env ]; then
    echo "Loading Cypress environment"
    . cypress.env
fi

cypress open
