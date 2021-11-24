SHELL := /bin/bash

all: \
	proto \
	go-lint \
	go-review \
	go-test \
	go-mod-tidy \
	readme-suites \
	prettier-format-readme \
	git-verify-nodiff

include tools/commitlint/rules.mk
include tools/git-verify-nodiff/rules.mk
include tools/golangci-lint/rules.mk
include tools/goreview/rules.mk
include tools/prettier/rules.mk
include tools/semantic-release/rules.mk
include tools/snippet/rules.mk

.PHONY: proto
proto:
	$(info [$@] building protos...)
	@make -C proto

.PHONY: readme-suites
readme-suites: $(snippet)
	$(info [$@] writing suites to README...)
	@go run ./cmd/doc | $(snippet) -M SUITES_SNIPPET -F README.md

.PHONY: prettier-format-readme
prettier-format-readme: $(prettier)
	$(info [$@] formatting README...)
	@$(prettier) --write 'README.md' --loglevel warn

.PHONY: go-test
go-test:
	$(info [$@] running Go tests...)
	@go test -count 1 -cover -race ./...

.PHONY: go-mod-tidy
go-mod-tidy:
	$(info [$@] tidying Go module files...)
	@go mod tidy -v
