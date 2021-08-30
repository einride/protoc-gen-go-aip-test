package search

import (
	"github.com/einride/protoc-gen-go-aip-test/internal/ident"
	"github.com/einride/protoc-gen-go-aip-test/internal/onlyif"
	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
	"github.com/einride/protoc-gen-go-aip-test/internal/util"
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/protobuf/compiler/protogen"
)

var deleted = suite.Test{
	Name: "deleted",
	Doc: []string{
		"Method should not return deleted resources.",
	},

	OnlyIf: suite.OnlyIfs(
		onlyif.HasMethod(aipreflect.MethodTypeSearch),
		onlyif.HasMethod(aipreflect.MethodTypeDelete),
		onlyif.HasParent,
	),
	Generate: func(f *protogen.GeneratedFile, scope suite.Scope) error {
		searchMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeSearch)
		deleteMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeDelete)
		responseResources := aipreflect.GrammaticalName(scope.Resource.GetPlural()).UpperCamelCase()
		f.P("const deleteCount = 5")
		f.P("for i := 0; i < deleteCount; i++ {")
		util.MethodDelete{
			Method:   deleteMethod,
			Resource: scope.Resource,
			Name:     "parentMsgs[i].Name",
		}.Generate(f, "_", "err", ":=")
		f.P(ident.AssertNilError, "(t, err)")
		f.P("}")
		util.MethodSearch{
			Resource: scope.Resource,
			Method:   searchMethod,
			Parent:   "parent",
			PageSize: "9999",
		}.Generate(f, "response", "err", ":=")
		f.P(ident.AssertNilError, "(t, err)")
		f.P(ident.AssertDeepEqual, "(")
		f.P("t,")
		f.P("parentMsgs[deleteCount:],")
		f.P("response.", responseResources, ",")
		f.P(ident.CmpoptsSortSlices, "(func(a,b *", scope.Message.GoIdent, ") bool {")
		f.P("return a.Name < b.Name")
		f.P("}),")
		f.P(ident.ProtocmpTransform, "(),")
		f.P(")")
		return nil
	},
}
