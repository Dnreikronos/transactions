build:
	@go build -o bin/transactions cmd/main.go

run: build
	@./bin/transactions cmd/main.go

test:
	@go test -v ./...