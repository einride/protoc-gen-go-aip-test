package search

import (
	"github.com/einride/protoc-gen-go-aip-test/internal/ident"
	"github.com/einride/protoc-gen-go-aip-test/internal/onlyif"
	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
	"github.com/einride/protoc-gen-go-aip-test/internal/util"
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/protobuf/compiler/protogen"
)

var lastPage = suite.Test{
	Name: "last page",
	Doc: []string{
		"If there are no more resources, next_page_token should not be set.",
	},

	OnlyIf: suite.OnlyIfs(
		onlyif.HasMethod(aipreflect.MethodTypeSearch),
		onlyif.HasParent,
	),
	Generate: func(f *protogen.GeneratedFile, scope suite.Scope) error {
		searchMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeSearch)
		util.MethodSearch{
			Resource: scope.Resource,
			Method:   searchMethod,
			Parent:   "parent",
			PageSize: "resourcesCount",
		}.Generate(f, "response", "err", ":=")
		f.P(ident.AssertNilError, "(t, err)")
		f.P(ident.AssertEqual, "(t, \"\", response.NextPageToken)")
		return nil
	},
}
