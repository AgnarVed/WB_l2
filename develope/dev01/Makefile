.PHONY: build run test

build:
		go mod tidy
		make test
		go build -o dev01 ./...
		make run
run:
		./dev01
test:
		golint ./...
		go vet ./...
		go test ./...

.DEFAULT_GOAL := build