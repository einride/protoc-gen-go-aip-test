SHELL := /bin/bash

all: \
	buf-lint \
	buf-generate

include tools/buf/rules.mk
include tools/protoc-gen-go/rules.mk
include tools/protoc-gen-go-grpc/rules.mk


.PHONY: buf-lint
buf-lint: $(buf)
	$(info [$@] linting protobuf schemas...)
	@$(buf) lint

protoc_gen_go_aip_test := ./bin/protoc-gen-go-aip-test
export PATH := $(dir $(abspath $(protoc_gen_go_aip_test))):$(PATH)

.PHONY: $(protoc_gen_go_aip_test)
$(protoc_gen_go_aip_test):
	$(info [$@] building binary...)
	go build -o $@ ../

.PHONY: buf-generate
buf-generate: $(buf) $(protoc_gen_go_aip_test) $(protoc_gen_go) $(protoc_gen_go_grpc)
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
