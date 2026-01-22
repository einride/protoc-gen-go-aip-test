package create

import (
	"strconv"

	"github.com/einride/protoc-gen-go-aip-test/internal/ident"
	"github.com/einride/protoc-gen-go-aip-test/internal/onlyif"
	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
	"github.com/einride/protoc-gen-go-aip-test/internal/util"
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protopath"
	"google.golang.org/protobuf/reflect/protoreflect"
)

//nolint:gochecknoglobals
var resourceReferences = suite.Test{
	Name: "resource references",
	Doc: []string{
		"The method should fail with InvalidArgument if the resource has any",
		"resource references and they are invalid.",
	},

	OnlyIf: suite.OnlyIfs(
		onlyif.HasMethod(aipreflect.MethodTypeCreate),
		onlyif.HasMutableResourceReferences,
	),
	Generate: func(f *protogen.GeneratedFile, scope suite.Scope, apiMode util.APIMode) error {
		createMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeCreate)
		util.RangeMutableResourceReferences(
			scope.Message.Desc,
			func(p protopath.Path, field protoreflect.FieldDescriptor, _ *annotations.ResourceReference) {
				if field.ContainingOneof() != nil {
					// resource references that are also one-ofs are
					// tricky to test this way
					return
				}
				// strip root step
				p = p[1:]
				containerPath := p[:len(p)-1]
				fieldPath := p[len(p)-1]
				isTopLevel := len(containerPath) == 0

				f.P("t.Run(", strconv.Quote(p.String()), ", func(t *", ident.TestingT, ") {")
				f.P(ident.FixtureMaybeSkip, "(t)")
				if util.HasParent(scope.Resource) {
					f.P("parent := ", ident.FixtureNextParent, "(t, false)")
				}
				if util.HasParent(scope.Resource) {
					f.P("msg := fx.Create(parent)")
				} else {
					f.P("msg := fx.Create()")
				}
				if isTopLevel {
					f.P("container := msg")
				} else {
					f.P("container := msg.", util.PathChainGet(containerPath))
				}
				f.P("if container == nil {")
				f.P("t.Skip(\"not reachable\")")
				f.P("}")

				fieldName := string(fieldPath.FieldDescriptor().Name())
				if field.IsList() {
					f.P(util.FieldSet("container", fieldName, "[]string{\"invalid resource name\"}", apiMode))
				} else {
					f.P(util.FieldSet("container", fieldName, "\"invalid resource name\"", apiMode))
				}
				util.MethodCreate{
					Resource: scope.Resource,
					Method:   createMethod,
					Parent:   "parent",
					Message:  "msg",
				}.Generate(f, "req", "_", "err", ":=", apiMode)
				f.P(
					ident.AssertEqual,
					"(t, ",
					ident.Codes(codes.InvalidArgument),
					", ",
					ident.StatusCode,
					"(err), err)",
				)
				f.P("})")
			},
		)
		return nil
	},
}
