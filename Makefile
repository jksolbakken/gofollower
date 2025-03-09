.PHONY: all compile test check

all: compile test check

compile: 
	go build -o bin/gofollower cmd/gofollower/*.go

test:
	go test -v ./...

check:
	go run honnef.co/go/tools/cmd/staticcheck ./...
	go run golang.org/x/vuln/cmd/govulncheck ./...