package update

import (
	"github.com/einride/protoc-gen-go-aip-test/internal/ident"
	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
	"github.com/einride/protoc-gen-go-aip-test/internal/util"
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/compiler/protogen"
)

var notFound = suite.Test{
	Name: "not found",
	Doc: []string{
		"Method should fail with NotFound if the resource does not exist.",
	},

	OnlyIf: func(scope suite.Scope) bool {
		updateMethod, ok := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeUpdate)
		return ok && !util.ReturnsLRO(updateMethod.Desc)
	},
	Generate: func(f *protogen.GeneratedFile, scope suite.Scope) error {
		updateMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeUpdate)
		util.MethodUpdate{
			Resource: scope.Resource,
			Method:   updateMethod,
			Parent:   "parent",
			// appending to the resource name ensures it is valid
			Name: "created.Name + \"notfound\"",
		}.Generate(f, "_", "err", ":=")
		f.P(ident.AssertEqual, "(t, ", ident.Codes(codes.NotFound), ",", ident.StatusCode, "(err), err)")
		return nil
	},
}
