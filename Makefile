# run server
run:
	go run main.go

# run server without a db
run-memory:
	go run main.go -use-memory

# run all tests with verbose output
test:
	go test ./... -v

# lint entire project
lint:
	golint ./...
