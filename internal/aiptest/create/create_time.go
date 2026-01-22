package create

import (
	"github.com/einride/protoc-gen-go-aip-test/internal/ident"
	"github.com/einride/protoc-gen-go-aip-test/internal/onlyif"
	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
	"github.com/einride/protoc-gen-go-aip-test/internal/util"
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/protobuf/compiler/protogen"
)

//nolint:gochecknoglobals
var createTime = suite.Test{
	Name: "create time",
	Doc: []string{
		"Field create_time should be populated when the resource is created.",
	},

	OnlyIf: suite.OnlyIfs(
		onlyif.HasMethod(aipreflect.MethodTypeCreate),
		onlyif.MethodNotLRO(aipreflect.MethodTypeCreate),
		onlyif.HasField("create_time"),
	),

	Generate: func(f *protogen.GeneratedFile, scope suite.Scope, apiMode util.APIMode) error {
		createMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeCreate)
		if util.HasParent(scope.Resource) {
			f.P("parent := ", ident.FixtureNextParent, "(t, false)")
		}
		f.P("beforeCreate := ", ident.TimeNow, "()")
		util.MethodCreate{
			Resource: scope.Resource,
			Method:   createMethod,
			Parent:   "parent",
		}.Generate(f, "req", "msg", "err", ":=", apiMode)
		f.P(ident.AssertNilError, "(t, err)")
		f.P(ident.AssertCheck, "(t, ", util.FieldGet("msg", "CreateTime", apiMode), " != nil)")
		f.P(ident.AssertCheck, "(t, !", util.FieldGet("msg", "CreateTime", apiMode), ".AsTime().IsZero())")
		f.P(
			ident.AssertCheck,
			"(t, ",
			util.FieldGet("msg", "CreateTime", apiMode),
			".AsTime().After(beforeCreate), ",
			"\"",
			util.FieldGet("msg", "CreateTime", apiMode),
			"(%v) is not after beforeCreate (%v)\", ",
			util.FieldGet("msg", "CreateTime", apiMode),
			".AsTime(), beforeCreate)",
		)
		return nil
	},
}
