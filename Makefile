build:
	@go build ./cmd/api

run:
	@go run ./cmd/api

lint:
	golangci-lint run

test:
	@go test -v ./...