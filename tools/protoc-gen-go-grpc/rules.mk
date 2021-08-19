grpc_go_cwd := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
protoc_gen_go_grpc := $(grpc_go_cwd)/bin/protoc-gen-go_grpc
export PATH := $(dir $(protoc_gen_go_grpc)):$(PATH)

$(protoc_gen_go_grpc): $(grpc_go_cwd)/../../go.mod
	$(info [protoc-gen-go-grpc] building binary..)
	@cd $(grpc_go_cwd)/../.. && go build -o $@ google.golang.org/grpc/cmd/protoc-gen-go-grpc
