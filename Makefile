# run server
run:
	go run main.go

# run server without a db
run-memory:
	go run main.go -use-memory

# run all tests with verbose output
test:
	go test ./... -v

# run all unit and integration tests
test-integration:
	docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit

# lint entire project
lint:
	$(shell go env GOPATH)/bin/golangci-lint run

# docker
run-services:
	docker-compose up --build

build-image:
	docker build -t go-todo .

run-image:
	docker run -p 3000:3000 go-todo

run-image-memory:
	docker run -p 3000:3000 go-todo -use-memory
