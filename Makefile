default: run
.PHONY: run

# Vue.js
run:
	cd frontend && npm run dev

# Golang
go-run:
	cd apigolang && go run cmd/main.go -config ./../configLocal.env

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