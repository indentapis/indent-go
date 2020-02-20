.PHONY: build test

# Go
GOPKG := go.indent.com/indent-go

build: pkg/auditv1
	go build -v $(GOPKG)/$<

test:
	go test -v ./...
