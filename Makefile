default: run
.PHONY: run

run:
	cd frontend && npm run dev

# migrate:	
	

run-docker-compose:
	docker compose --env-file configForDocker.env up --build -d

go-lint:
	cd apigolang && golangci-lint run ./... -c .golangci.yml

go-memory-check:
	fieldalignment ./...

go-memory-fix:
	fieldalignment -fix ./...