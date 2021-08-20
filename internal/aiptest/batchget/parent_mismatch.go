package batchget

import (
	"github.com/einride/protoc-gen-go-aip-test/internal/ident"
	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
	"github.com/einride/protoc-gen-go-aip-test/internal/util"
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/compiler/protogen"
)

var parentMismatch = suite.Test{
	Name: "parent mismatch",
	Doc: []string{
		"If a caller sets the \"parent\", and the parent collection in the name of any resource",
		"being retrieved does not match, the request must fail.",
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
			Parent:   "fx.peekNextParent(t)",
			Names:    []string{"created00.Name"},
		}.Generate(f, "_", "err", ":=")
		f.P(ident.AssertEqual, "(t, ", ident.Codes(codes.InvalidArgument), ",", ident.StatusCode, "(err), err)")
		return nil
	},
}
