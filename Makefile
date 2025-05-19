.PHONY: build run clean test cover vet fmt lint mocks test-file test-one deps docker-up docker-down docker-build

build:
	go build -o bin/app ./cmd/app

run:
	go run ./cmd/app

clean:
	go clean
	rm -rf bin coverage.out

test:
	go test ./... -v

cover:
	go test ./... -coverprofile=coverage.out

cover-html:
	go tool cover -html=coverage.out

vet:
	go vet ./...

fmt:
	go fmt ./...

lint:
	golangci-lint run ./...

mocks:
	mockery --all --output=mocks --with-expecter

test-file:
	go test -v ./services/post_service_test.go

test-one:
	go test -v ./services -run ^TestNewPostService$

deps:
	go mod tidy

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

docker-build:
	docker-compose build