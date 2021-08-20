package batchget

import (
	"strconv"

	"github.com/einride/protoc-gen-go-aip-test/internal/ident"
	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
	"github.com/einride/protoc-gen-go-aip-test/internal/util"
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/compiler/protogen"
)

var parentInvalid = suite.Test{
	Name: "invalid parent",
	Doc: []string{
		"Method should fail with InvalidArgument if provided parent is invalid.",
	},

	OnlyIf: func(scope suite.Scope) bool {
		batchGetMethod, hasBatchGet := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeBatchGet)
		return hasBatchGet &&
			util.HasParent(scope.Resource) &&
			!util.IsAlternativeBatchGet(batchGetMethod.Desc)
	},
	Generate: func(f *protogen.GeneratedFile, scope suite.Scope) error {
		batchGetMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeBatchGet)
		util.MethodBatchGet{
			Resource: scope.Resource,
			Method:   batchGetMethod,
			Parent:   strconv.Quote("invalid resource name"),
		}.Generate(f, "_", "err", ":=")
		f.P(ident.AssertEqual, "(t, ", ident.Codes(codes.InvalidArgument), ", ", ident.StatusCode, "(err), err)")
		return nil
	},
}
