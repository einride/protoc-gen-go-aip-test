package search

import (
	"github.com/einride/protoc-gen-go-aip-test/internal/ident"
	"github.com/einride/protoc-gen-go-aip-test/internal/onlyif"
	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
	"github.com/einride/protoc-gen-go-aip-test/internal/util"
	"github.com/stoewer/go-strcase"
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/protobuf/compiler/protogen"
)

var oneByOne = suite.Test{
	Name: "one by one",
	Doc: []string{
		"Searching resource one by one should eventually return all resources.",
	},

	OnlyIf: suite.OnlyIfs(
		onlyif.HasMethod(aipreflect.MethodTypeSearch),
		onlyif.HasParent,
	),
	Generate: func(f *protogen.GeneratedFile, scope suite.Scope) error {
		searchMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeSearch)
		responseResources := strcase.UpperCamelCase(string(util.FindResourceField(
			searchMethod.Output.Desc,
			scope.Resource,
		).Name()))
		f.P("msgs := make([]*", scope.Message.GoIdent, ", 0, resourcesCount)")
		f.P("var nextPageToken string")
		f.P("for {")
		util.MethodSearch{
			Resource:  scope.Resource,
			Method:    searchMethod,
			Parent:    "parent",
			PageToken: "nextPageToken",
			PageSize:  "1",
		}.Generate(f, "response", "err", ":=")
		f.P(ident.AssertNilError, "(t, err)")
		f.P(ident.AssertEqual, "(t, 1, len(response.", responseResources, "))")
		f.P("msgs = append(msgs, response.", responseResources, "...)")
		f.P("nextPageToken = response.NextPageToken")
		f.P("if nextPageToken == \"\" {")
		f.P("break")
		f.P("}")
		f.P("}")
		f.P(ident.AssertDeepEqual, "(")
		f.P("t,")
		f.P("parentMsgs,")
		f.P("msgs,")
		f.P(ident.CmpoptsSortSlices, "(func(a,b *", scope.Message.GoIdent, ") bool {")
		f.P("return a.Name < b.Name")
		f.P("}),")
		f.P(ident.ProtocmpTransform, "(),")
		f.P(")")
		return nil
	},
}
