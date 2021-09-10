VERSION ?= dev

test:
	go test ./...
.PHONY:test

build: test
	go build -ldflags "-X main.Version=${VERSION}"
.PHONY:build

install: test
	go install -ldflags "-X main.Version=${VERSION}"
.PHONY:install
