package list

import (
	"github.com/einride/protoc-gen-go-aip-test/internal/ident"
	"github.com/einride/protoc-gen-go-aip-test/internal/onlyif"
	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
	"google.golang.org/protobuf/compiler/protogen"
)

// Suite of List tests.
//
//nolint:gochecknoglobals
var Suite = suite.Suite{
	Name: "List",
	Tests: []suite.Test{
		invalidParent,
		invalidPageToken,
		negativePageSize,
	},
	TestGroups: []suite.TestGroup{
		withResourcesGroup,
		withBigResourcesGroup,
	},
}

//nolint:gochecknoglobals
var withResourcesGroup = suite.TestGroup{
	OnlyIf: suite.OnlyIfs(
		onlyif.HasParent,
	),
	GenerateBefore: func(f *protogen.GeneratedFile, scope suite.Scope) error {
		f.P("const resourcesCount = 15")
		f.P("parent := ", ident.FixtureNextParent, "(t, true)")
		f.P("parentMsgs := make([]*", scope.Message.GoIdent, ", resourcesCount)")
		f.P("for i := 0; i < resourcesCount; i++ {")
		f.P("parentMsgs[i] = fx.create(t, parent)")
		f.P("}")
		f.P()
		return nil
	},
	Tests: []suite.Test{
		isolation,
		lastPage,
		morePages,
		oneByOne,
		deleted,
	},
}

//nolint:gochecknoglobals
var withBigResourcesGroup = suite.TestGroup{
	OnlyIf: suite.OnlyIfs(
		onlyif.HasParent,
	),
	GenerateBefore: func(f *protogen.GeneratedFile, scope suite.Scope) error {
		f.P("const resourcesCount = 101")
		f.P("parent := ", ident.FixtureNextParent, "(t, true)")
		f.P("parentMsgs := make([]*", scope.Message.GoIdent, ", resourcesCount)")
		f.P("for i := 0; i < resourcesCount; i++ {")
		f.P("parentMsgs[i] = fx.create(t, parent)")
		f.P("}")
		f.P()
		return nil
	},
	Tests: []suite.Test{
		pageSizeZero,
	},
}
