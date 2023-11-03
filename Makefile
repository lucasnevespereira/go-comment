.PHONY: clean build lint

APP_NAME=gocomment

fmt:
	gofmt -s -l -w .

lint: fmt
	golangci-lint run

build:
	go build -o $(APP_NAME)

