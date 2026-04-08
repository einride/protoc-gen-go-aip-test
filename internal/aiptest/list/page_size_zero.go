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
		"When listing resource with page size zero the service should use a default value.",
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
		util.MethodList{
			Resource: scope.Resource,
			Method:   listMethod,
			Parent:   "parent",
			PageSize: "0",
		}.Generate(f, "response", "err", ":=")
		f.P(ident.AssertNilError, "(t, err)")
		f.P("// Server should use a default page size and return at least some results")
		f.P(
			ident.AssertCheck,
			"(t, len(response.",
			responseResources,
			") > 0, \"expected server to return at least 1 resource with page_size=0\")",
		)
		return nil
	},
}
