include .env

MAKEFILE_DIR := $(dir $(realpath $(firstword $(MAKEFILE_LIST))))
DC = docker compose -f $(MAKEFILE_DIR)compose.yml
DM = docker run --rm -it -v ./infra/mysql/migration:/migration --network gin-sample-network  migrate/migrate -path=/migration/ -database 'mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DATABASE)'

# for Help
%:
	@:

# for Local
#------------------------------------------------------------------------------
.PHONY: run-dev
run-dev:
	air -c .air.toml

# for Docker
#------------------------------------------------------------------------------

.PHONY: up, down
up:
	$(DC) up -d
down:
	$(DC) down

.PHONY: log-api, log-mysql
log-api:
	$(eval ID := $(shell docker ps | rg 'gin-sample-api-1' | cut -d' ' -f1))
	docker logs -f $(ID)
log-mysql:
	$(eval ID := $(shell docker ps | rg 'gin-sample-mysql-1' | cut -d' ' -f1))
	docker logs -f $(ID)

.PHONY: exec-api, exec-mysql
exec-api:
	$(DC) exec api /bin/bash || true
exec-mysql:
	$(DC) exec mysql /bin/bash || true

# Lint
#
.PHONY: lint, lint-fix
lint:
	docker run --pull always --rm -v $$(pwd):/app:ro -w /app golangci/golangci-lint:v2.5.0-alpine golangci-lint run --timeout=5m
lint-fix:
	docker run --pull always --rm -v $$(pwd):/app -w /app golangci/golangci-lint:v2.5.0-alpine golangci-lint run --fix --timeout=5m

# for migrate
#------------------------------------------------------------------------------
.PHONY: migrate-create, migrate-up, migrate-down, migrate-drop
migrate-create:
	$(DM) create -ext sql -dir /migration $(filter-out $@,$(MAKECMDGOALS))
migrate-up:
	$(DM) up
migrate-down:
	$(DM) down 1
migrate-drop:
	$(DM) drop
migrate-reset:
	$(DM) drop
	$(DM) up

# for generate model
#------------------------------------------------------------------------------
.PHONY: generate-model
generate-model:
	$(DC) exec api /bin/bash -c "gentool -db mysql -dsn '$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DATABASE)' -fieldNullable -fieldSignable -fieldWithIndexTag -fieldWithTypeTag -modelPkgName model -onlyModel -outPath ./infra/mysql/model/"
