SHELL := /bin/bash

all: \
	buf-lint \
	buf-generate \
	go-lint \
	go-review \
	buf-generate \
	go-test \
	go-mod-tidy \
	git-verify-nodiff

include tools/buf/rules.mk
include tools/commitlint/rules.mk
include tools/git-verify-nodiff/rules.mk
include tools/golangci-lint/rules.mk
include tools/goreview/rules.mk
include tools/protoc-gen-go/rules.mk
include tools/protoc-gen-go-grpc/rules.mk
include tools/semantic-release/rules.mk

.PHONY: go-test
go-test:
	$(info [$@] running Go tests...)
	@go test -count 1 -cover -race ./...

.PHONY: go-mod-tidy
go-mod-tidy:
	$(info [$@] tidying Go module files...)
	@go mod tidy -v

.PHONY: buf-lint
buf-lint: $(buf)
	$(info [$@] linting protobuf schemas...)
	@$(buf) lint

protoc_gen_go_aip_test := ./bin/protoc-gen-go-aip-test
export PATH := $(dir $(abspath $(protoc_gen_go_aip_test))):$(PATH)

.PHONY: $(protoc_gen_go_aip_test)
$(protoc_gen_go_aip_test):
	$(info [$@] building binary...)
	@go build -o $@ .

.PHONY: buf-generate
buf-generate: $(buf) $(protoc_gen_go_aip_test) $(protoc_gen_go) $(protoc_gen_go_grpc)
	$(info [$@] generating protobuf stubs...)
	@rm -rf proto/gen
	@$(buf) generate --path proto/src/einride
