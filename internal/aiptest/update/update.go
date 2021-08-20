package update

import (
	"github.com/einride/protoc-gen-go-aip-test/internal/ident"
	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
	"github.com/einride/protoc-gen-go-aip-test/internal/util"
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/protobuf/compiler/protogen"
)

var Suite = suite.Suite{
	Name: "Update",
	Tests: []suite.Test{
		missingName,
		invalidName,
	},
	TestGroups: []suite.TestGroup{
		withResourceGroup,
	},
}

var withResourceGroup = suite.TestGroup{
	OnlyIf: func(scope suite.Scope) bool {
		createMethod, hasCreate := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeCreate)
		return hasCreate && !util.ReturnsLRO(createMethod.Desc)
	},
	GenerateBefore: func(f *protogen.GeneratedFile, scope suite.Scope) error {
		createMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeCreate)
		if util.HasParent(scope.Resource) {
			f.P("parent := ", ident.FixtureNextParent, "(t, false)")
		}
		util.MethodCreate{
			Resource: scope.Resource,
			Method:   createMethod,
			Parent:   "parent",
		}.Generate(f, "created", "err", ":=")
		f.P(ident.AssertNilError, "(t, err)")
		return nil
	},
	Tests: []suite.Test{
		updateTime,
		notFound,
		persisted,
		invalidUpdateMask,
		requiredFields,
		// TODO: add test for supplying wildcard as name
		// TODO: add test for etags
	},
}
