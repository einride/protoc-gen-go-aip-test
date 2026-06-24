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

//nolint:gochecknoglobals
var allExists = suite.Test{
	Name: "all exists",
	Doc: []string{
		"Resources should be returned without errors if they exist.",
	},

	OnlyIf: suite.OnlyIfs(
		onlyif.HasMethod(aipreflect.MethodTypeBatchGet),
		onlyif.BatchMethodNotAlternative(aipreflect.MethodTypeBatchGet),
	),
	Generate: func(f *protogen.GeneratedFile, scope suite.Scope, apiMode util.APIMode) error {
		batchGetMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeBatchGet)
		responseResources := strcase.UpperCamelCase(string(util.FindResourceField(
			batchGetMethod.Output.Desc,
			scope.Resource,
		).Name()))
		names := []string{"created00", "created01", "created02"}
		getters := make([]string, 0, len(names))
		for _, name := range names {
			getters = append(getters, util.FieldGet(name, "Name", apiMode))
		}
		util.MethodBatchGet{
			Resource: scope.Resource,
			Method:   batchGetMethod,
			Parent:   "parent",
			Names:    getters,
		}.Generate(f, "req", "response", "err", ":=", apiMode)
		f.P(ident.AssertNilError, "(t, err)")
		f.P(ident.AssertDeepEqual, "(")
		f.P("t,")
		f.P("[]*", scope.Message.GoIdent, "{")
		for _, name := range names {
			f.P(name + ",")
		}
		f.P("},")
		f.P(util.FieldGet("response", responseResources, apiMode), ",")
		f.P(ident.ProtocmpTransform, "(),")
		f.P(")")
		return nil
	},
}
