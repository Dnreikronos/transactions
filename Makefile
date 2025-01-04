build:
	@go build -o bin/transactions cmd/main.go

run: build
	@./bin/transactions cmd/main.go

testHandler:
	@go test -v ./tests/handler_test.go
testWorker:
	@go test -v ./tests/worker_test.go
testQueue:
	@go test -v ./tests/queue_test.go