package create

import (
	"strconv"
	"strings"

	"github.com/einride/protoc-gen-go-aip-test/internal/ident"
	"github.com/einride/protoc-gen-go-aip-test/internal/onlyif"
	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
	"github.com/einride/protoc-gen-go-aip-test/internal/util"
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/grpc/codes"
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

//nolint:gochecknoglobals
var invalidUserSettableID = suite.Test{
	Name: "invalid user settable id",
	Doc: []string{
		"Method should fail with InvalidArgument if the user settable id doesn't",
		"conform to RFC-1034, see [doc](https://google.aip.dev/122#resource-id-segments).",
	},
	OnlyIf: suite.OnlyIfs(
		onlyif.HasMethod(aipreflect.MethodTypeCreate),
		onlyif.MethodNotLRO(aipreflect.MethodTypeCreate),
		onlyif.HasUserSettableID,
	),
	Generate: func(f *protogen.GeneratedFile, scope suite.Scope) error {
		f.P("for _, tt := range []struct {")
		f.P("name string")
		f.P("id string")
		f.P("} {")
		for _, tt := range []struct {
			name string
			id   string
		}{
			{
				name: "start with digit",
				id:   "0foo",
			},
			{
				name: "start with hyphen",
				id:   "-foo",
			},
			{
				name: "start with non ascii letter",
				id:   "öfoo",
			},
			{
				name: "contains non ascii letter",
				id:   "fooöbar",
			},
			{
				name: "contains upper case ascii letter",
				id:   "fooBar",
			},
			{
				name: "ends with hyphen",
				id:   "foo-",
			},
			{
				name: "ends with non ascii",
				id:   "fooö",
			},
			{
				name: "too short",
				id:   "foo",
			},
			{
				name: "too long",
				id:   "f" + strings.Repeat("o", 63),
			},
		} {
			f.P("{")
			f.P("name: ", strconv.Quote(tt.name), ",")
			f.P("id: ", strconv.Quote(tt.id), ",")
			f.P("},")
		}
		f.P("} {")
		f.P("t.Run(tt.name, func(t *testing.T) {")
		f.P(ident.FixtureMaybeSkip, "(t)")
		createMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeCreate)
		if util.HasParent(scope.Resource) {
			f.P("parent := ", ident.FixtureNextParent, "(t, false)")
		}
		util.MethodCreate{
			Resource:       scope.Resource,
			Method:         createMethod,
			Parent:         "parent",
			UserSettableID: "tt.id",
		}.Generate(f, "_", "err", ":=")
		f.P(ident.AssertEqual, "(t, ", ident.Codes(codes.InvalidArgument), ", ", ident.StatusCode, "(err), err)")
		f.P("})")
		f.P("}")
		return nil
	},
}
