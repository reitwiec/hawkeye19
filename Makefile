include .env
export $(shell sed 's/=.*//' .env)

SHELL := /bin/bash

# Start server
serve:
	go run api/main.go

# Generate keys
keys:
	go run scripts/generate_keys.go

# Migrate
migrate:
	go run api/main.go -m

# Drop all tables
drop_all:
	@echo 'Dropping all tables in $(HAWK_DB_NAME)'
	@mysql --user=$(HAWK_DB_USER) --password=$(HAWK_DB_PASSWORD) --database=$(HAWK_DB_NAME) < ./scripts/drop.sql
