.PHONY: dev test clear run

clear:
	@rm -rf ./bin/

dev:
	@go run ./cmd

test:
	@go test ./...

build: clear
	@go build -o bin/main ./cmd

run: build
	@bin/main

