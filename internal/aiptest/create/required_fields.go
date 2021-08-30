package create

import (
	"strconv"

	"github.com/einride/protoc-gen-go-aip-test/internal/ident"
	"github.com/einride/protoc-gen-go-aip-test/internal/onlyif"
	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
	"github.com/einride/protoc-gen-go-aip-test/internal/util"
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protopath"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var requiredFields = suite.Test{
	Name: "required fields",
	Doc: []string{
		"The method should fail with InvalidArgument if the resource has any",
		"required fields and they are not provided.",
	},

	OnlyIf: suite.OnlyIfs(
		onlyif.HasMethod(aipreflect.MethodTypeCreate),
		onlyif.HasRequiredFields,
	),
	Generate: func(f *protogen.GeneratedFile, scope suite.Scope) error {
		createMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeCreate)
		util.RangeRequiredFields(scope.Message.Desc, func(p protopath.Path, field protoreflect.FieldDescriptor) {
			// strip root step
			p = p[1:]
			containerPath := p[:len(p)-1]
			fieldPath := p[len(p)-1]
			isTopLevel := len(containerPath) == 0

			f.P("t.Run(", strconv.Quote(p.String()), ", func(t *", ident.TestingT, ") {")
			f.P("fx.maybeSkip(t)")
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
			f.P("fd := container.ProtoReflect().Descriptor().Fields().ByName(", strconv.Quote(fieldName), ")")
			f.P("container.ProtoReflect().Clear(fd)")
			util.MethodCreate{
				Resource: scope.Resource,
				Method:   createMethod,
				Parent:   "parent",
				Message:  "msg",
			}.Generate(f, "_", "err", ":=")
			f.P(ident.AssertEqual, "(t, ", ident.Codes(codes.InvalidArgument), ", ", ident.StatusCode, "(err), err)")
			f.P("})")
		})
		return nil
	},
}
