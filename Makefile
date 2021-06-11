.PHONY:
.SILENT:
.DEFAULT_GOAL := run

build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/main.go

run: build
	docker-compose up --remove-orphans app

test:
	go test --short -coverprofile=cover.out -v ./...
	make test.coverage

create-migration: # запуск спараметром: make create-migration NAME=test_migration
	migrate create -ext sql -dir schema/ -seq $(NAME)

# Testing vars example
export TEST_DB_URI=mongodb://localhost:27019
export TEST_DB_NAME=test
export TEST_CONTAINER_NAME=test_db
# Testing command example
test.integration:
	docker run --rm -d -p 27019:27017 --name $$TEST_CONTAINER_NAME -e MONGODB_DATABASE=$$TEST_DB_NAME mongo:4.4-bionic

	GIN_MODE=release go test -v ./tests/
	docker stop $$TEST_CONTAINER_NAME

test.coverage:
	go tool cover -func=cover.out

swag:
	swag init -g cmd/main.go

lint:
	golangci-lint run