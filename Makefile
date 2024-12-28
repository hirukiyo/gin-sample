include .env

MAKEFILE_DIR := $(dir $(realpath $(firstword $(MAKEFILE_LIST))))
DC = docker compose -f $(MAKEFILE_DIR)compose.yml

# for Local
.PHONY: run-dev
run-dev:
	air -c .air.toml

# for Docker
.PHONY: up, down
up:
	$(DC) up -d

down:
	$(DC) down

.PHONY: log-api, log-mysql
log-api:
	$(eval ID := $(shell docker ps | rg 'gin-api-1' | cut -d' ' -f1))
	docker logs -f $(ID)

log-mysql:
	$(eval ID := $(shell docker ps | rg 'gin-mysql-1' | cut -d' ' -f1))
	docker logs -f $(ID)

.PHONY: exec-api, exec-mysql
exec-api:
	$(DC) exec api /bin/bash

exec-mysql:
	$(DC) exec mysql /bin/bash
