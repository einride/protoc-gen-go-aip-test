package create

import (
	"github.com/einride/protoc-gen-go-aip-test/internal/ident"
	"github.com/einride/protoc-gen-go-aip-test/internal/onlyif"
	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
	"github.com/einride/protoc-gen-go-aip-test/internal/util"
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/protobuf/compiler/protogen"
)

var createTime = suite.Test{
	Name: "create time",
	Doc: []string{
		"Field create_time should be populated when the resource is created.",
	},

	OnlyIf: suite.OnlyIfs(
		onlyif.HasMethod(aipreflect.MethodTypeCreate),
		onlyif.MethodNotLRO(aipreflect.MethodTypeCreate),
	),
	Generate: func(f *protogen.GeneratedFile, scope suite.Scope) error {
		createMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeCreate)
		if util.HasParent(scope.Resource) {
			f.P("parent := ", ident.FixtureNextParent, "(t, false)")
		}
		util.MethodCreate{
			Resource: scope.Resource,
			Method:   createMethod,
			Parent:   "parent",
		}.Generate(f, "msg", "err", ":=")
		f.P(ident.AssertNilError, "(t, err)")
		f.P(ident.AssertCheck, "(t, ", ident.TimeSince, "(msg.CreateTime.AsTime()) < ", ident.TimeSecond, ")")
		return nil
	},
}