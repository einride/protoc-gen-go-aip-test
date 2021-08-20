package update

import (
	"strconv"

	"github.com/einride/protoc-gen-go-aip-test/internal/ident"
	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
	"github.com/einride/protoc-gen-go-aip-test/internal/util"
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/compiler/protogen"
)

var missingName = suite.Test{
	Name: "missing name",
	Doc: []string{
		"Method should fail with InvalidArgument if no name is provided.",
	},

	OnlyIf: func(scope suite.Scope) bool {
		_, ok := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeUpdate)
		return ok
	},
	Generate: func(f *protogen.GeneratedFile, scope suite.Scope) error {
		updateMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeUpdate)
		if util.HasParent(scope.Resource) {
			f.P("parent := ", ident.FixtureNextParent, "(t, false)")
		}
		util.MethodUpdate{
			Resource: scope.Resource,
			Method:   updateMethod,
			Parent:   "parent",
			Name:     strconv.Quote(""),
		}.Generate(f, "_", "err", ":=")
		f.P(ident.AssertEqual, "(t, ", ident.Codes(codes.InvalidArgument), ",", ident.StatusCode, "(err), err)")
		return nil
	},
}
