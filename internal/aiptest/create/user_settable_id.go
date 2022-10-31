package create

import (
	"strconv"

	"github.com/einride/protoc-gen-go-aip-test/internal/ident"
	"github.com/einride/protoc-gen-go-aip-test/internal/onlyif"
	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
	"github.com/einride/protoc-gen-go-aip-test/internal/util"
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/protobuf/compiler/protogen"
)

//nolint:gochecknoglobals
var userSettableID = suite.Test{
	Name: "user settable id",
	Doc: []string{
		"If method support user settable IDs, when set the resource should",
		"be returned with the provided ID.",
	},

	OnlyIf: suite.OnlyIfs(
		onlyif.HasMethod(aipreflect.MethodTypeCreate),
		onlyif.MethodNotLRO(aipreflect.MethodTypeCreate),
		onlyif.HasUserSettableID,
	),
	Generate: func(f *protogen.GeneratedFile, scope suite.Scope) error {
		createMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeCreate)
		if util.HasParent(scope.Resource) {
			f.P("parent := ", ident.FixtureNextParent, "(t, false)")
		}
		util.MethodCreate{
			Resource:       scope.Resource,
			Method:         createMethod,
			Parent:         "parent",
			UserSettableID: strconv.Quote("usersetid"),
		}.Generate(f, "msg", "err", ":=")
		f.P(ident.AssertNilError, "(t, err)")
		f.P(ident.AssertCheck, "(t, ", ident.StringsHasSuffix, "(msg.GetName(), ", strconv.Quote("usersetid"), "))")
		return nil
	},
}
