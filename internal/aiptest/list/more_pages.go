package list

import (
	"github.com/einride/protoc-gen-go-aip-test/internal/ident"
	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
	"github.com/einride/protoc-gen-go-aip-test/internal/util"
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/protobuf/compiler/protogen"
)

var morePages = suite.Test{
	Name: "more pages",
	Doc: []string{
		"If there are more resources, next_page_token should be set.",
	},

	OnlyIf: func(scope suite.Scope) bool {
		_, hasList := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeList)
		return hasList && util.HasParent(scope.Resource)
	},
	Generate: func(f *protogen.GeneratedFile, scope suite.Scope) error {
		listMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeList)
		util.MethodList{
			Resource: scope.Resource,
			Method:   listMethod,
			Parent:   "parent",
			PageSize: "resourcesCount",
		}.Generate(f, "response", "err", ":=")
		f.P(ident.AssertNilError, "(t, err)")
		f.P(ident.AssertCheck, "(t, response.NextPageToken != \"\")")
		return nil
	},
}
