package get

import (
	"github.com/einride/protoc-gen-go-aip-test/internal/ident"
	"github.com/einride/protoc-gen-go-aip-test/internal/onlyif"
	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
	"github.com/einride/protoc-gen-go-aip-test/internal/util"
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/protobuf/compiler/protogen"
)

// nolint: gochecknoglobals
var exists = suite.Test{
	Name: "exists",
	Doc: []string{
		"Resource should be returned without errors if it exists.",
	},

	OnlyIf: suite.OnlyIfs(
		onlyif.HasMethod(aipreflect.MethodTypeGet),
	),
	Generate: func(f *protogen.GeneratedFile, scope suite.Scope) error {
		getMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeGet)

		if util.HasParent(scope.Resource) {
			f.P("parent := ", ident.FixtureNextParent, "(t, false)")
			f.P("created := fx.create(t, parent)")
		} else {
			f.P("created := fx.create(t)")
		}
		util.MethodGet{
			Resource: scope.Resource,
			Method:   getMethod,
			// appending to the resource name ensures it is valid
			Name: "created.Name",
		}.Generate(f, "msg", "err", ":=")
		f.P(ident.AssertNilError, "(t, err)")
		f.P(ident.AssertDeepEqual, "(t, msg, created, ", ident.ProtocmpTransform, "())")
		return nil
	},
}
