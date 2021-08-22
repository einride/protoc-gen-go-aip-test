package update

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
		"The updated resource should be persisted and reachable with Get.",
	},

	OnlyIf: func(scope suite.Scope) bool {
		updateMethod, hasUpdate := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeUpdate)
		createMethod, hasCreate := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeCreate)
		_, hasGet := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeGet)
		return hasUpdate && !util.ReturnsLRO(updateMethod.Desc) &&
			hasCreate && !util.ReturnsLRO(createMethod.Desc) &&
			hasGet
	},
	Generate: func(f *protogen.GeneratedFile, scope suite.Scope) error {
		updateMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeUpdate)
		getMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeGet)
		createMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeCreate)
		if util.HasParent(scope.Resource) {
			f.P("parent := ", ident.FixtureNextParent, "(t, false)")
		}
		util.MethodCreate{
			Resource: scope.Resource,
			Method:   createMethod,
			Parent:   "parent",
		}.Generate(f, "created", "err", ":=")
		f.P(ident.AssertNilError, "(t, err)")
		util.MethodUpdate{
			Resource: scope.Resource,
			Method:   updateMethod,
			Msg:      "created",
		}.Generate(f, "updated", "err", ":=")
		f.P(ident.AssertNilError, "(t, err)")
		util.MethodGet{
			Resource: scope.Resource,
			Method:   getMethod,
			Name:     "updated.Name",
		}.Generate(f, "persisted", "err", ":=")
		f.P(ident.AssertNilError, "(t, err)")
		f.P(ident.AssertDeepEqual, "(t, updated, persisted, ", ident.ProtocmpTransform, "())")
		return nil
	},
}
