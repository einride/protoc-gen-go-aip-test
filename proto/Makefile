SHELL := /bin/bash

all: \
	buf-lint \
	buf-generate

include .tools/buf/rules.mk

.build/protoc-gen-go: ../go.mod
	$(info [$@] building binary...)
	@go build -o $@ google.golang.org/protobuf/cmd/protoc-gen-go

.build/protoc-gen-go-grpc: $(abspath $(lastword $(MAKEFILE_LIST)))
	$(info [$@] building binary...)
	@GOBIN=$(abspath $(dir $@)) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0

.PHONY: .build/protoc-gen-go-aip-test
.build/protoc-gen-go-aip-test:
	$(info [$@] building binary...)
	@cd .. && go build -o $(abspath $@) .

.PHONY: buf-lint
buf-lint: $(buf)
	$(info [$@] linting protobuf schemas...)
	@$(buf) lint

protoc_plugins := \
	.build/protoc-gen-go \
	.build/protoc-gen-go-grpc \
	.build/protoc-gen-go-aip-test

.PHONY: buf-generate
buf-generate: $(buf) $(protoc_plugins)
	$(info [$@] generating protobuf stubs...)
	@rm -rf proto/gen
	@$(buf) generate --path einride
	@$(buf) generate buf.build/googleapis/googleapis \
		--template buf.gen.googleapis.yaml \
		--path google/area120/tables/v1alpha1 \
		--path google/cloud/aiplatform/v1 \
		--path google/cloud/gsuiteaddons/v1 \
		--path google/cloud/scheduler/v1 \
		--path google/pubsub/v1 \
		--path google/spanner
