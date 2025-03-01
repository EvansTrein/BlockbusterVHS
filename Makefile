default: run
.PHONY: run

run:
	cd frontend && npm run dev

# migrate:	
	

run-docker-compose:
	docker compose --env-file configForDocker.env up --build -d

go-run:
	cd apigolang && go run cmd/main.go -config ./../configLocal.env

go-lint:
	cd apigolang && golangci-lint run ./... -c .golangci.yml

go-format:
	cd apigolang && go fmt ./...

go-memory-check:
	fieldalignment ./...

go-memory-fix:
	fieldalignment -fix ./...