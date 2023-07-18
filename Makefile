DOCKERCOMPOSE ?= docker-compose

GO ?= go
GOFMT ?= gofmt "-s"
GOFILES := $(shell find . -name "*.go")
GOMODULES := $(shell go list ./...)

# clean & build

clean:
	@rm -r build

build:
	$(GO) build -o build/program/app cmd/httpd/main.go

# run & dev

run:
	make build
	./build/program/app

dev:
	$(GO) run cmd/httpd/main.go

# fmt

fmt:
	$(GOFMT) -w $(GOFILES)

# testing

test:
	$(GO) clean -testcache
	$(GO) mod tidy
	$(GO) test -cover $(GOMODULES)

# docker

compose:
	$(DOCKERCOMPOSE) up -d

# phony

.PHONY: clean build run dev fmt test compose