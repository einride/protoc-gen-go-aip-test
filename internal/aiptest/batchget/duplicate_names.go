package batchget

import (
	"github.com/einride/protoc-gen-go-aip-test/internal/ident"
	"github.com/einride/protoc-gen-go-aip-test/internal/onlyif"
	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
	"github.com/einride/protoc-gen-go-aip-test/internal/util"
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/protobuf/compiler/protogen"
)

var duplicateNames = suite.Test{
	Name: "duplicate names",
	Doc: []string{
		"If a caller provides duplicate names, the service should return",
		"duplicate resources.",
	},

	OnlyIf: suite.OnlyIfs(
		onlyif.HasMethod(aipreflect.MethodTypeBatchGet),
		onlyif.BatchMethodNotAlternative(aipreflect.MethodTypeBatchGet),
	),
	Generate: func(f *protogen.GeneratedFile, scope suite.Scope) error {
		batchGetMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeBatchGet)
		responseResources := aipreflect.GrammaticalName(scope.Resource.GetPlural()).UpperCamelCase()
		util.MethodBatchGet{
			Resource: scope.Resource,
			Method:   batchGetMethod,
			Parent:   "parent",
			Names:    []string{"created00.Name", "created00.Name"},
		}.Generate(f, "response", "err", ":=")
		f.P(ident.AssertNilError, "(t, err)")
		f.P(ident.AssertDeepEqual, "(")
		f.P("t,")
		f.P("[]*", scope.Message.GoIdent, "{")
		f.P("created00,")
		f.P("created00,")
		f.P("},")
		f.P("response.", responseResources, ",")
		f.P(ident.ProtocmpTransform, "(),")
		f.P(")")
		return nil
	},
}
