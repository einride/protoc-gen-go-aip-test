package batchget

import (
	"github.com/einride/protoc-gen-go-aip-test/internal/ident"
	"github.com/einride/protoc-gen-go-aip-test/internal/onlyif"
	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
	"github.com/einride/protoc-gen-go-aip-test/internal/util"
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/compiler/protogen"
)

var namesMissing = suite.Test{
	Name: "names missing",
	Doc: []string{
		"Method should fail with InvalidArgument if no names are provided.",
	},

	OnlyIf: suite.OnlyIfs(
		onlyif.HasMethod(aipreflect.MethodTypeBatchGet),
		onlyif.BatchMethodNotAlternative(aipreflect.MethodTypeBatchGet),
	),
	Generate: func(f *protogen.GeneratedFile, scope suite.Scope) error {
		batchGetMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeBatchGet)
		if util.HasParent(scope.Resource) {
			f.P("parent := ", ident.FixtureNextParent, "(t, false)")
		}
		util.MethodBatchGet{
			Resource: scope.Resource,
			Method:   batchGetMethod,
			Parent:   "parent",
			Names:    nil,
		}.Generate(f, "_", "err", ":=")
		f.P(ident.AssertEqual, "(t, ", ident.Codes(codes.InvalidArgument), ", ", ident.StatusCode, "(err), err)")
		return nil
	},
}
