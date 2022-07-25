ifeq ($(OS),Windows_NT)
    GOOS := windows
else
    UNAME_S := $(shell uname -s)
    ifeq ($(UNAME_S),Linux)
        GOOS := linux
    endif
    ifeq ($(UNAME_S),Darwin)
        GOOS := darwin
    endif
endif

CI_COMMIT_SHORT_SHA ?= $(shell git rev-parse --short HEAD)
CGO_ENABLED = 0
GOARCH = amd64
LDFLAGS = -ldflags "-X main.shaCommit=${CI_COMMIT_SHORT_SHA}"
GO = $(shell which go)
GO_BUILD = GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=$(CGO_ENABLED) $(GO) build $(LDFLAGS)

################################################### tests

.PHONY: test
test:
	CGO_ENABLED=1 $(GO) test -p 1 -race -short ./pkg/... ./internal/...

################################################### run

.PHONY: run
run:
	go run ./cmd/app

################################################### utils

.PHONY: format
format:
	$(GO) fmt ./...

.PHONY: install
install:
	$(GO) mod download

################################################### migrations

.PHONY: create-migration
create-migration:
	go run ./cmd/migration ./cmd/app/migrations

################################################### build

.PHONY: generate
generate:
	$(GO) generate ./...

################################################### email templates

.PHONY: compile-email-templates
compile-email-templates:
	yarn mjml ./cmd/app/templates/action/template.mjml -o ./cmd/app/templates/action/html.jet
	yarn mjml ./cmd/app/templates/notification/template.mjml -o ./cmd/app/templates/notification/html.jet

################################################### osrm - data

.PHONY: lint-revive
lint-revive:
	revive -config ./revive.toml -exclude ./internal/graph/generated ./...

include Makefile.build.mk
include Makefile.docker.mk
