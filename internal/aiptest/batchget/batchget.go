package batchget

import (
	"github.com/einride/protoc-gen-go-aip-test/internal/ident"
	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
	"github.com/einride/protoc-gen-go-aip-test/internal/util"
	"google.golang.org/protobuf/compiler/protogen"
)

// Suite of BatchGet tests.
//
//nolint:gochecknoglobals
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

//nolint:gochecknoglobals
var withResourcesGroup = suite.TestGroup{
	GenerateBefore: func(f *protogen.GeneratedFile, scope suite.Scope) error {
		if util.HasParent(scope.Resource) {
			f.P("parent := ", ident.FixtureNextParent, "(t, false)")
		}
		for i := 0; i < 3; i++ {
			if util.HasParent(scope.Resource) {
				f.P("created0", i, " := fx.create(t, parent)")
			} else {
				f.P("created0", i, " := fx.create(t)")
			}
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
