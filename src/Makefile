.PHONY: test test-services test-file test-run lint

test:
	go test ./... -v

test-services:
	go test ./services -v

test-file:
	cd services && go test -v post_service_test.go

test-run:
	cd services && go test -v -run ^TestNewPostService$

lint:
	golangci-lint run ./...