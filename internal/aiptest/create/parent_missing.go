package create

import (
	"strconv"

	"github.com/einride/protoc-gen-go-aip-test/internal/ident"
	"github.com/einride/protoc-gen-go-aip-test/internal/onlyif"
	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
	"github.com/einride/protoc-gen-go-aip-test/internal/util"
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/compiler/protogen"
)

// nolint: gochecknoglobals
var parentMissing = suite.Test{
	Name: "missing parent",
	Doc: []string{
		"Method should fail with InvalidArgument if no parent is provided.",
	},

	OnlyIf: suite.OnlyIfs(
		onlyif.HasMethod(aipreflect.MethodTypeCreate),
		onlyif.HasParent,
	),
	Generate: func(f *protogen.GeneratedFile, scope suite.Scope) error {
		createMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeCreate)
		util.MethodCreate{
			Resource: scope.Resource,
			Method:   createMethod,
			Message:  "fx.Create(fx.nextParent(t, false))",
			Parent:   strconv.Quote(""),
		}.Generate(f, "_", "err", ":=")
		f.P(ident.AssertEqual, "(t, ", ident.Codes(codes.InvalidArgument), ", ", ident.StatusCode, "(err), err)")
		return nil
	},
}
