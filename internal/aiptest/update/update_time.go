package update

import (
	"github.com/einride/protoc-gen-go-aip-test/internal/ident"
	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
	"github.com/einride/protoc-gen-go-aip-test/internal/util"
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/protobuf/compiler/protogen"
)

var updateTime = suite.Test{
	Name: "update time",
	Doc: []string{
		"Field update_time should be updated when the resource is updated.",
	},

	OnlyIf: func(scope suite.Scope) bool {
		updateMethod, hasUpdate := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeUpdate)
		createMethod, hasCreate := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeCreate)
		return hasUpdate && !util.ReturnsLRO(updateMethod.Desc) &&
			hasCreate && !util.ReturnsLRO(createMethod.Desc)
	},
	Generate: func(f *protogen.GeneratedFile, scope suite.Scope) error {
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
		updateMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeUpdate)
		util.MethodUpdate{
			Resource: scope.Resource,
			Method:   updateMethod,
			Msg:      "created",
		}.Generate(f, "updated", "err", ":=")
		f.P(ident.AssertNilError, "(t, err)")
		f.P(ident.AssertCheck, "(t, updated.UpdateTime.AsTime().After(created.UpdateTime.AsTime()))")
		return nil
	},
}
