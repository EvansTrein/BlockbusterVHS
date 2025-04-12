default: run
.PHONY: run

PATH_DB=sqlite://database/database.db
# PATH_DB=postgres://evans:evans@localhost:8012/postgres?sslmode=disable
FILE_MIGRATIONS =./migrations
MIGRATION_MODE=up

# Vue.js
run:
	cd frontend && npm run dev

# Golang
go-run:
	cd apigolang && go run cmd/main.go -config ./configLocal.env

go-migrate-up:	
	cd apigolang && go run cmd/migrator/migrator.go -mode $(MIGRATION_MODE) -storage-path $(PATH_DB) -migrations-path $(FILE_MIGRATIONS)

go-lint:
	cd apigolang && golangci-lint run ./... -c .golangci.yml

go-fmt:
	cd apigolang && go fmt ./...

go-memory-check:
	cd apigolang && fieldalignment ./...

go-memory-fix:
	cd apigolang && fieldalignment -fix ./...

go-mock:
	cd apigolang && go generate ./...

go-cover:
	cd apigolang && go test -cover ./...

go-cover-html:
	cd apigolang && go test -cover -coverprofile=coverage.out ./internal/users && go tool cover -html=coverage.out -o coverage.html



run-docker-compose:
	docker compose --env-file ./apigolang/configForDocker.env up --build -d