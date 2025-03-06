package list

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
var pageSizeZero = suite.Test{
	Name: "page size zero",
	Doc: []string{
		"Listing resource with page size zero should eventually return all resources.",
	},

	OnlyIf: suite.OnlyIfs(
		onlyif.HasMethod(aipreflect.MethodTypeList),
		onlyif.HasParent,
	),
	Generate: func(f *protogen.GeneratedFile, scope suite.Scope) error {
		listMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeList)
		responseResources := strcase.UpperCamelCase(string(util.FindResourceField(
			listMethod.Output.Desc,
			scope.Resource,
		).Name()))
		f.P("msgs := make([]*", scope.Message.GoIdent, ", 0, resourcesCount)")
		f.P("var nextPageToken string")
		f.P("for {")
		util.MethodList{
			Resource:  scope.Resource,
			Method:    listMethod,
			Parent:    "parent",
			PageSize:  "0",
			PageToken: "nextPageToken",
		}.Generate(f, "page", "err", ":=")
		f.P(ident.AssertNilError, "(t, err)")
		f.P("msgs = append(msgs, page." + responseResources + "...)")
		f.P("nextPageToken = page.NextPageToken")
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
