.PHONY: all gofollower test check

all: gofollower test check

gofollower:
	go build -o bin/follow cmd/gofollower/main.go

test:
	go test -v ./...

check:
	go run honnef.co/go/tools/cmd/staticcheck ./...
	go run golang.org/x/vuln/cmd/govulncheck ./...
	go run github.com/securego/gosec/v2/cmd/gosec@v2.22.10 --exclude-generated -terse ./...