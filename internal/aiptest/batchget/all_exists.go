package batchget

import (
	"github.com/einride/protoc-gen-go-aip-test/internal/ident"
	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
	"github.com/einride/protoc-gen-go-aip-test/internal/util"
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/protobuf/compiler/protogen"
)

var allExists = suite.Test{
	Name: "all exists",
	Doc: []string{
		"Resources should be returned without errors if they exist.",
	},

	OnlyIf: func(scope suite.Scope) bool {
		batchGetMethod, hasBatchGet := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeBatchGet)
		return hasBatchGet &&
			!util.IsAlternativeBatchGet(batchGetMethod.Desc)
	},
	Generate: func(f *protogen.GeneratedFile, scope suite.Scope) error {
		batchGetMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeBatchGet)
		responseResources := aipreflect.GrammaticalName(scope.Resource.GetPlural()).UpperCamelCase()
		util.MethodBatchGet{
			Resource: scope.Resource,
			Method:   batchGetMethod,
			Parent:   "parent",
			Names:    []string{"created00.Name", "created01.Name", "created02.Name"},
		}.Generate(f, "response", "err", ":=")
		f.P(ident.AssertNilError, "(t, err)")
		f.P(ident.AssertDeepEqual, "(")
		f.P("t,")
		f.P("[]*", scope.Message.GoIdent, "{")
		f.P("created00,")
		f.P("created01,")
		f.P("created02,")
		f.P("},")
		f.P("response.", responseResources, ",")
		f.P(ident.ProtocmpTransform, "(),")
		f.P(")")
		return nil
	},
}
