package search

import (
	"strconv"

	"github.com/einride/protoc-gen-go-aip-test/internal/ident"
	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
	"github.com/einride/protoc-gen-go-aip-test/internal/util"
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/compiler/protogen"
)

var invalidParent = suite.Test{
	Name: "invalid parent",
	Doc: []string{
		"Method should fail with InvalidArgument if provided parent is invalid.",
	},

	OnlyIf: func(scope suite.Scope) bool {
		_, hasSearch := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeSearch)
		return hasSearch && util.HasParent(scope.Resource)
	},
	Generate: func(f *protogen.GeneratedFile, scope suite.Scope) error {
		searchMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeSearch)
		util.MethodSearch{
			Resource: scope.Resource,
			Method:   searchMethod,
			Parent:   strconv.Quote("invalid resource name"),
		}.Generate(f, "_", "err", ":=")
		f.P(ident.AssertEqual, "(t, ", ident.Codes(codes.InvalidArgument), ", ", ident.StatusCode, "(err), err)")
		return nil
	},
}
