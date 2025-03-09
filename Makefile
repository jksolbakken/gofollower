.PHONY: all compile test

all: compile test

compile: 
	go build -o bin/gofollower cmd/gofollower/*.go

test:
	go test -v ./...

