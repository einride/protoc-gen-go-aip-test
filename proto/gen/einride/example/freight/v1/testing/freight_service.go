package examplefreightv1test

import (
	context "context"
	v1 "github.com/einride/protoc-gen-go-aiptest/proto/gen/einride/example/freight/v1"
)

type FreightService struct {
	// Context to use for running tests.
	Context context.Context

	// The service to test.
	Service v1.FreightServiceServer
}
