build:
	@go build ./cmd/api

run:
	@go run ./cmd/api

test:
	@go test -v ./...