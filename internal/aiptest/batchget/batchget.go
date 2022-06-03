package batchget

import (
	"fmt"

	"github.com/einride/protoc-gen-go-aip-test/internal/ident"
	"github.com/einride/protoc-gen-go-aip-test/internal/onlyif"
	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
	"github.com/einride/protoc-gen-go-aip-test/internal/util"
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/protobuf/compiler/protogen"
)

// Suite of BatchGet tests.
// nolint: gochecknoglobals
var Suite = suite.Suite{
	Name: "BatchGet",
	Tests: []suite.Test{
		parentInvalid,
		namesMissing,
		namesInvalid,
		wildcardName,
		// TODO: add test for supplying wildcard as parent
	},
	TestGroups: []suite.TestGroup{
		withResourcesGroup,
	},
}

// nolint: gochecknoglobals
var withResourcesGroup = suite.TestGroup{
	OnlyIf: suite.OnlyIfs(
		onlyif.HasMethod(aipreflect.MethodTypeCreate),
		onlyif.MethodNotLRO(aipreflect.MethodTypeCreate),
	),
	GenerateBefore: func(f *protogen.GeneratedFile, scope suite.Scope) error {
		createMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeCreate)
		if util.HasParent(scope.Resource) {
			f.P("parent := ", ident.FixtureNextParent, "(t, false)")
		}
		for i := 0; i < 3; i++ {
			util.MethodCreate{
				Resource: scope.Resource,
				Method:   createMethod,
				Parent:   "parent",
			}.Generate(f, fmt.Sprintf("created0%d", i), "err", ":=")
			f.P(ident.AssertNilError, "(t, err)")
		}
		return nil
	},
	Tests: []suite.Test{
		allExists,
		atomic,
		parentMismatch,
		ordered,
		duplicateNames,
	},
}
