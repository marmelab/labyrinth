#!/usr/bin/env bash

while ! nc -z postgres 5432 ; do sleep 1 ; done

php bin/console doctrine:migrations:migrate --no-interaction
