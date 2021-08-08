package plugin

import (
	"strconv"

	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/protobuf/compiler/protogen"
)

func (r *resourceGenerator) createTestCase() testCase {
	createMethodName, err := r.resource.InferMethodName(aipreflect.MethodTypeCreate)
	if err != nil {
		return disabledTestCase()
	}
	createMethod, ok := findMethod(r.service, createMethodName)
	if !ok {
		return disabledTestCase()
	}

	return newTestCase("Create", func(f *protogen.GeneratedFile) {
		testingT := f.QualifiedGoIdent(protogen.GoIdent{GoName: "T", GoImportPath: "testing"})
		timeSecond := f.QualifiedGoIdent(protogen.GoIdent{GoName: "Second", GoImportPath: "time"})
		timeSince := f.QualifiedGoIdent(protogen.GoIdent{GoName: "Since", GoImportPath: "time"})
		stringsHasSuffix := f.QualifiedGoIdent(protogen.GoIdent{GoName: "HasSuffix", GoImportPath: "strings"})
		assertCheck := f.QualifiedGoIdent(protogen.GoIdent{
			GoName:       "Check",
			GoImportPath: "gotest.tools/v3/assert",
		})
		assertEqual := f.QualifiedGoIdent(protogen.GoIdent{
			GoName:       "Equal",
			GoImportPath: "gotest.tools/v3/assert",
		})
		assertNilError := f.QualifiedGoIdent(protogen.GoIdent{
			GoName:       "NilError",
			GoImportPath: "gotest.tools/v3/assert",
		})
		statusCode := f.QualifiedGoIdent(protogen.GoIdent{
			GoName:       "Code",
			GoImportPath: "google.golang.org/grpc/status",
		})
		codesInvalidArgument := f.QualifiedGoIdent(protogen.GoIdent{
			GoName:       "InvalidArgument",
			GoImportPath: "google.golang.org/grpc/codes",
		})
		codesAlreadyExists := f.QualifiedGoIdent(protogen.GoIdent{
			GoName:       "AlreadyExists",
			GoImportPath: "google.golang.org/grpc/codes",
		})

		f.P("// Standard methods: Create")
		f.P("// https://google.aip.dev/133")

		if hasParent(r.resource) {
			f.P()
			f.P("parent := fx.nextParent(t, false)")
		}

		if hasParent(r.resource) {
			f.P()
			f.P("// Method should fail with InvalidArgument if no parent is provided.")
			f.P("t.Run(\"missing parent\", func(t *", testingT, ") {")
			f.P("fx.maybeSkip(t)")
			m := methodCreate{
				resource: r.resource,
				method:   createMethod,
				parent:   strconv.Quote(""),
			}
			m.Generate(f, "_", "err", ":=")
			f.P(assertEqual, "(t, ", codesInvalidArgument, ",", statusCode, "(err), err)")
			f.P("})")
		}

		if hasParent(r.resource) {
			f.P()
			f.P("// Method should fail with InvalidArgument is provided parent is not valid.")
			f.P("t.Run(\"invalid parent\", func(t *", testingT, ") {")
			f.P("fx.maybeSkip(t)")
			m := methodCreate{
				resource: r.resource,
				method:   createMethod,
				parent:   strconv.Quote("invalid resource name"),
			}
			m.Generate(f, "_", "err", ":=")
			f.P(assertEqual, "(t, ", codesInvalidArgument, ",", statusCode, "(err), err)")
			f.P("})")
		}

		f.P()
		f.P("// Field create_time should be populated when the resource is created.")
		f.P("t.Run(\"create time\", func(t *", testingT, ") {")
		f.P("fx.maybeSkip(t)")
		m := methodCreate{
			resource: r.resource,
			method:   createMethod,
			parent:   "parent",
		}
		m.Generate(f, "msg", "err", ":=")
		f.P(assertNilError, "(t, err)")
		f.P(assertCheck, "(t, ", timeSince, "(msg.CreateTime.AsTime()) < ", timeSecond, ")")
		f.P("})")

		if hasUserSettableID(r.resource, createMethod.Desc) {
			f.P()
			f.P("// If method support user settable IDs, when set the resource should")
			f.P("// returned with the provided ID.")
			f.P("t.Run(\"user settable id\", func(t *", testingT, ") {")
			f.P("fx.maybeSkip(t)")
			m := methodCreate{
				resource:       r.resource,
				method:         createMethod,
				parent:         "parent",
				userSettableID: strconv.Quote("usersetid"),
			}
			m.Generate(f, "msg", "err", ":=")
			f.P(assertNilError, "(t, err)")
			f.P(assertCheck, "(t, ", stringsHasSuffix, "(msg.GetName(), ", strconv.Quote("usersetid"), "))")
			f.P("})")
		}

		if hasUserSettableID(r.resource, createMethod.Desc) {
			f.P()
			f.P("// If method support user settable IDs and the same ID is reused")
			f.P("// the method should return AlreadyExists.")
			f.P("t.Run(\"already exists\", func(t *", testingT, ") {")
			f.P("fx.maybeSkip(t)")
			m := methodCreate{
				resource:       r.resource,
				method:         createMethod,
				parent:         "parent",
				userSettableID: strconv.Quote("alreadyexists"),
			}
			m.Generate(f, "_", "err", ":=")
			f.P(assertNilError, "(t, err)")
			m.Generate(f, "_", "err", "=")
			f.P(assertEqual, "(t, ", codesAlreadyExists, ",", statusCode, "(err), err)")
			f.P("})")
		}
	})
}
