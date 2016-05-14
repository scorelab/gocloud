.PHONY: all test clean build install

GOFLAGS ?= $(GOFLAGS:)

all: install test


build:
	@godep go build $(GOFLAGS) ./...

install:
	@godep go install $(GOFLAGS) ./...

test: install
	@godep go install $(GOFLAGS) ./...

bench: install
	@godep go test -run=NONE -bench=. $(GOFLAGS) ./...

clean:
	@godep go clean $(GOFLAGS) -i ./...

init:
	@godep save ./...