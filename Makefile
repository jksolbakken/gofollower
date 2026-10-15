.PHONY: all gofollower test check

all: gofollower test check

gofollower:
	go build -o bin/follow cmd/gofollower/main.go

test:
	go test -v ./...

check:
	go tool staticcheck ./...
	go tool govulncheck ./...
	go tool gosec --exclude-generated -terse ./...
