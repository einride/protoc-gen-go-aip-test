package create

import (
	"github.com/einride/protoc-gen-go-aip-test/internal/ident"
	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
	"github.com/einride/protoc-gen-go-aip-test/internal/util"
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/protobuf/compiler/protogen"
)

var persisted = suite.Test{
	Name: "persisted",
	Doc: []string{
		"The created resource should be persisted and reachable with Get.",
	},

	OnlyIf: func(scope suite.Scope) bool {
		createMethod, hasCreate := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeCreate)
		_, hasGet := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeGet)
		return hasCreate && !util.ReturnsLRO(createMethod.Desc) && hasGet
	},
	Generate: func(f *protogen.GeneratedFile, scope suite.Scope) error {
		createMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeCreate)
		getMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeGet)
		if util.HasParent(scope.Resource) {
			f.P("parent := ", ident.FixtureNextParent, "(t, false)")
		}
		util.MethodCreate{
			Resource: scope.Resource,
			Method:   createMethod,
			Parent:   "parent",
		}.Generate(f, "msg", "err", ":=")
		f.P(ident.AssertNilError, "(t, err)")
		util.MethodGet{
			Resource: scope.Resource,
			Method:   getMethod,
			Name:     "msg.Name",
		}.Generate(f, "persisted", "err", ":=")
		f.P(ident.AssertNilError, "(t, err)")
		f.P(ident.AssertDeepEqual, "(t, msg, persisted, ", ident.ProtocmpTransform, "())")
		return nil
	},
}
