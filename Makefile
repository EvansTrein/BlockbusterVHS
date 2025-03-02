default: run
.PHONY: run

PATH_DB=./internal/storages/database.db
FILE_MIGRATIONS =./migrations

# Vue.js
run:
	cd frontend && npm run dev

# Golang
go-run:
	cd apigolang && go run cmd/main.go -config ./../configLocal.env

go-migrate-up:	
	cd apigolang && go run cmd/migrator/migrator.go -mode up -storage-path $(PATH_DB) -migrations-path $(FILE_MIGRATIONS)

go-lint:
	cd apigolang && golangci-lint run ./... -c .golangci.yml

go-format:
	cd apigolang && go fmt ./...

go-memory-check:
	cd apigolang && fieldalignment ./...

go-memory-fix:
	cd apigolang && fieldalignment -fix ./...



run-docker-compose:
	docker compose --env-file configForDocker.env up --build -d