build: ## build
	@go build -o bin/cockroach main.go

test: ## test
	@go test -v ./...

run: build ## run
	@./bin/cockroach

setting-db: ## setting infra
	@docker-compose -f infra/docker-compose.yml up -d

down-db: ## down mysql docker container
	@docker-compose -f infra/docker-compose.yml down

migrate:
	 @go run ./cockroach/migrations/cockroachMigrate.go

help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'