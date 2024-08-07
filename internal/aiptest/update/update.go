package update

import (
	"github.com/einride/protoc-gen-go-aip-test/internal/ident"
	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
	"github.com/einride/protoc-gen-go-aip-test/internal/util"
	"google.golang.org/protobuf/compiler/protogen"
)

// Suite for the Updaet method.
//
//nolint:gochecknoglobals
var Suite = suite.Suite{
	Name: "Update",
	Tests: []suite.Test{
		missingName,
		invalidName,
		updateTime,
		persisted,
		preserveCreateTime,
		etagMismatch,
		etagUpdated,
	},
	TestGroups: []suite.TestGroup{
		withResourceGroup,
	},
}

//nolint:gochecknoglobals
var withResourceGroup = suite.TestGroup{
	GenerateBefore: func(f *protogen.GeneratedFile, scope suite.Scope) error {
		if util.HasParent(scope.Resource) {
			f.P("parent := ", ident.FixtureNextParent, "(t, false)")
			f.P("created := fx.create(t, parent)")
		} else {
			f.P("created := fx.create(t)")
		}
		return nil
	},
	Tests: []suite.Test{
		notFound,
		invalidUpdateMask,
		requiredFields,
		// TODO: add test for supplying wildcard as name
	},
}
