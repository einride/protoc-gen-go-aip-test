package batchget

import (
	"github.com/einride/protoc-gen-go-aip-test/internal/ident"
	"github.com/einride/protoc-gen-go-aip-test/internal/onlyif"
	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
	"github.com/einride/protoc-gen-go-aip-test/internal/util"
	"github.com/stoewer/go-strcase"
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/protobuf/compiler/protogen"
)

//nolint: gochecknoglobals
var allExists = suite.Test{
	Name: "all exists",
	Doc: []string{
		"Resources should be returned without errors if they exist.",
	},

	OnlyIf: suite.OnlyIfs(
		onlyif.HasMethod(aipreflect.MethodTypeBatchGet),
		onlyif.BatchMethodNotAlternative(aipreflect.MethodTypeBatchGet),
	),
	Generate: func(f *protogen.GeneratedFile, scope suite.Scope) error {
		batchGetMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeBatchGet)
		responseResources := strcase.UpperCamelCase(string(util.FindResourceField(
			batchGetMethod.Output.Desc,
			scope.Resource,
		).Name()))
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
