package plugin

import (
	"strconv"

	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/protobuf/compiler/protogen"
)

func (r *resourceGenerator) getTestCase() testCase {
	getMethod, ok := r.standardMethod(aipreflect.MethodTypeGet)
	if !ok {
		return disabledTestCase()
	}
	createMethod, ok := r.standardMethod(aipreflect.MethodTypeCreate)
	if !ok {
		return disabledTestCase()
	}
	// TODO: support LROs for create.
	if returnsLRO(createMethod.Desc) {
		return disabledTestCase()
	}

	return newTestCase("Get", func(f *protogen.GeneratedFile) {
		testingT := f.QualifiedGoIdent(protogen.GoIdent{GoName: "T", GoImportPath: "testing"})
		assertEqual := f.QualifiedGoIdent(protogen.GoIdent{
			GoName:       "Equal",
			GoImportPath: "gotest.tools/v3/assert",
		})
		assertDeepEqual := f.QualifiedGoIdent(protogen.GoIdent{
			GoName:       "DeepEqual",
			GoImportPath: "gotest.tools/v3/assert",
		})
		assertNilError := f.QualifiedGoIdent(protogen.GoIdent{
			GoName:       "NilError",
			GoImportPath: "gotest.tools/v3/assert",
		})
		protocmpTransform := f.QualifiedGoIdent(protogen.GoIdent{
			GoName:       "Transform",
			GoImportPath: "google.golang.org/protobuf/testing/protocmp",
		})
		statusCode := f.QualifiedGoIdent(protogen.GoIdent{
			GoName:       "Code",
			GoImportPath: "google.golang.org/grpc/status",
		})
		codesInvalidArgument := f.QualifiedGoIdent(protogen.GoIdent{
			GoName:       "InvalidArgument",
			GoImportPath: "google.golang.org/grpc/codes",
		})
		codesNotFound := f.QualifiedGoIdent(protogen.GoIdent{
			GoName:       "NotFound",
			GoImportPath: "google.golang.org/grpc/codes",
		})

		f.P("// Standard methods: Get")
		f.P("// https://google.aip.dev/131")

		if hasParent(r.resource) {
			f.P()
			f.P("parent := fx.nextParent(t, false)")
		}
		methodCreate{
			resource: r.resource,
			method:   createMethod,
			parent:   "parent",
		}.Generate(f, "created00", "err", ":=")
		f.P(assertNilError, "(t, err)")

		f.P()
		f.P("// Method should fail with InvalidArgument if no name is provided.")
		f.P("t.Run(\"missing name\", func(t *", testingT, ") {")
		f.P("fx.maybeSkip(t)")
		methodGet{
			resource: r.resource,
			method:   getMethod,
			name:     strconv.Quote(""),
		}.Generate(f, "_", "err", ":=")
		f.P(assertEqual, "(t, ", codesInvalidArgument, ",", statusCode, "(err), err)")
		f.P("})")

		f.P()
		f.P("// Method should fail with InvalidArgument is provided name is not valid.")
		f.P("t.Run(\"invalid name\", func(t *", testingT, ") {")
		f.P("fx.maybeSkip(t)")
		methodGet{
			resource: r.resource,
			method:   getMethod,
			name:     strconv.Quote("invalid resource name"),
		}.Generate(f, "_", "err", ":=")
		f.P(assertEqual, "(t, ", codesInvalidArgument, ",", statusCode, "(err), err)")
		f.P("})")

		f.P()
		f.P("// Resource should be returned without errors if it exists.")
		f.P("t.Run(\"exists\", func(t *", testingT, ") {")
		f.P("fx.maybeSkip(t)")
		methodGet{
			resource: r.resource,
			method:   getMethod,
			name:     "created00.Name",
		}.Generate(f, "msg", "err", ":=")
		f.P(assertNilError, "(t, err)")
		f.P(assertDeepEqual, "(t, msg, created00, ", protocmpTransform, "())")
		f.P("})")

		f.P()
		f.P("// Method should fail with NotFound if the resource does not exist.")
		f.P("t.Run(\"not found\", func(t *", testingT, ") {")
		f.P("fx.maybeSkip(t)")
		methodGet{
			resource: r.resource,
			method:   getMethod,
			// appending to the resource name ensures it is valid
			name: "created00.Name + \"notfound\"",
		}.Generate(f, "_", "err", ":=")
		f.P(assertEqual, "(t, ", codesNotFound, ",", statusCode, "(err), err)")
		f.P("})")

		// TODO: add test for supplying wildcard as name

		f.P("_ = ", codesNotFound)
		f.P("_ = ", protocmpTransform)
	})
}
