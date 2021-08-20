package list

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
		_, hasList := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeList)
		return hasList && util.HasParent(scope.Resource)
	},
	Generate: func(f *protogen.GeneratedFile, scope suite.Scope) error {
		listMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeList)
		util.MethodList{
			Resource: scope.Resource,
			Method:   listMethod,
			Parent:   strconv.Quote("invalid resource name"),
		}.Generate(f, "_", "err", ":=")
		f.P(ident.AssertEqual, "(t, ", ident.Codes(codes.InvalidArgument), ", ", ident.StatusCode, "(err), err)")
		return nil
	},
}
