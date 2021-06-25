.PHONY:
.SILENT:
.DEFAULT_GOAL := run

build:
	go mod download && go build cmd/main.go

test:
	go test --short -coverprofile=cover.out -v ./...
	make test.coverage

create-migration:
	migrate create -ext sql -dir ./schema -seq init

test.coverage:
	go tool cover -func=cover.out

swag:
	swag init -g cmd/main.go

lint:
	golangci-lint run