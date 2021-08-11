package plugin

import (
	"fmt"
	"strconv"

	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/protobuf/compiler/protogen"
)

func (r *resourceGenerator) batchGetTestCase() testCase {
	batchGetMethod, ok := r.standardMethod(aipreflect.MethodTypeBatchGet)
	if !ok {
		return disabledTestCase()
	}
	createMethod, ok := r.standardMethod(aipreflect.MethodTypeCreate)
	if !ok {
		return disabledTestCase()
	}
	// TODO: support alternative btach get
	if isAlternativeBatchGet(batchGetMethod.Desc) {
		return disabledTestCase()
	}
	// TODO: support LROs for create.
	if returnsLRO(createMethod.Desc) {
		return disabledTestCase()
	}
	responseResources := aipreflect.GrammaticalName(r.resource.GetPlural()).UpperCamelCase()

	return newTestCase("BatchGet", func(f *protogen.GeneratedFile) {
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

		f.P("// Batch methods: Get")
		f.P("// https://google.aip.dev/231")

		if hasParent(r.resource) {
			f.P()
			f.P("parent := fx.nextParent(t, false)")
		}
		for i := 0; i < 3; i++ {
			methodCreate{
				resource: r.resource,
				method:   createMethod,
				parent:   "parent",
			}.Generate(f, fmt.Sprintf("created0%d", i), "err", ":=")
			f.P(assertNilError, "(t, err)")
		}

		if hasParent(r.resource) {
			f.P()
			f.P("// Method should fail with InvalidArgument if provided parent is not valid.")
			f.P("t.Run(\"invalid parent\", func(t *", testingT, ") {")
			f.P("fx.maybeSkip(t)")
			methodBatchGet{
				resource: r.resource,
				method:   batchGetMethod,
				parent:   strconv.Quote("invalid resource name"),
				names:    []string{"created00.Name"},
			}.Generate(f, "_", "err", ":=")
			f.P(assertEqual, "(t, ", codesInvalidArgument, ",", statusCode, "(err), err)")
			f.P("})")
		}

		f.P()
		f.P("// Method should fail with InvalidArgument if no names are provided.")
		f.P("t.Run(\"no names\", func(t *", testingT, ") {")
		f.P("fx.maybeSkip(t)")
		methodBatchGet{
			resource: r.resource,
			method:   batchGetMethod,
			parent:   "parent",
			names:    nil,
		}.Generate(f, "_", "err", ":=")
		f.P(assertEqual, "(t, ", codesInvalidArgument, ",", statusCode, "(err), err)")
		f.P("})")

		f.P()
		f.P("// Method should fail with InvalidArgument if a provided name is not valid.")
		f.P("t.Run(\"invalid name\", func(t *", testingT, ") {")
		f.P("fx.maybeSkip(t)")
		methodBatchGet{
			resource: r.resource,
			method:   batchGetMethod,
			parent:   "parent",
			names:    []string{strconv.Quote("invalid resource name")},
		}.Generate(f, "_", "err", ":=")
		f.P(assertEqual, "(t, ", codesInvalidArgument, ",", statusCode, "(err), err)")
		f.P("})")

		f.P()
		f.P("// Resources should be returned without errors if they exist.")
		f.P("t.Run(\"all exists\", func(t *", testingT, ") {")
		f.P("fx.maybeSkip(t)")
		methodBatchGet{
			resource: r.resource,
			method:   batchGetMethod,
			parent:   "parent",
			names:    []string{"created00.Name", "created01.Name", "created02.Name"},
		}.Generate(f, "response", "err", ":=")
		f.P(assertNilError, "(t, err)")
		f.P(assertDeepEqual, "(")
		f.P("t,")
		f.P("[]*", r.message.GoIdent, "{")
		f.P("created00,")
		f.P("created01,")
		f.P("created02,")
		f.P("},")
		f.P("response.", responseResources, ",")
		f.P(protocmpTransform, "(),")
		f.P(")")
		f.P("})")

		f.P()
		f.P("// The method must be atomic; it must fail for all resources")
		f.P("// or succeed for all resources (no partial success). ")
		f.P("t.Run(\"atomic\", func(t *", testingT, ") {")
		f.P("fx.maybeSkip(t)")
		methodBatchGet{
			resource: r.resource,
			method:   batchGetMethod,
			parent:   "parent",
			names: []string{
				"created00.Name",
				// appending to the resource name ensures it is valid
				"created01.Name + \"notfound\"",
				"created02.Name",
			},
		}.Generate(f, "_", "err", ":=")
		f.P(assertEqual, "(t, ", codesNotFound, ", ", statusCode, "(err), err)")
		f.P("})")

		if hasParent(r.resource) {
			f.P()
			f.P("// If a caller sets the \"parent\", and the parent collection in the name of any resource")
			f.P("// being retrieved does not match, the request must fail.")
			f.P("t.Run(\"parent mismatch\", func(t *", testingT, ") {")
			f.P("fx.maybeSkip(t)")
			methodBatchGet{
				resource: r.resource,
				method:   batchGetMethod,
				parent:   "fx.peekNextParent(t)",
				names:    []string{"created00.Name"},
			}.Generate(f, "_", "err", ":=")
			f.P(assertEqual, "(t, ", codesInvalidArgument, ",", statusCode, "(err), err)")
			f.P("})")
			f.P()
		}

		f.P()
		f.P("// The order of resources in the response must be the same as the names in the request.")
		f.P("t.Run(\"ordered\", func(t *", testingT, ") {")
		f.P("fx.maybeSkip(t)")
		f.P("for _, order := range [][]*", r.message.GoIdent, "{")
		f.P("{created00, created01, created02},")
		f.P("{created01, created00, created02},")
		f.P("{created02, created01, created00},")
		f.P("} {")
		methodBatchGet{
			resource: r.resource,
			method:   batchGetMethod,
			parent:   "parent",
			names:    []string{"order[0].GetName()", "order[1].GetName()", "order[2].GetName()"},
		}.Generate(f, "response", "err", ":=")
		f.P(assertNilError, "(t, err)")
		f.P(assertDeepEqual, "(t, order, response.", responseResources, ",", protocmpTransform, "())")
		f.P("}")
		f.P("})")

		f.P()
		f.P("// If a caller provides duplicate names, the service should return")
		f.P("// duplicate resources.")
		f.P("t.Run(\"duplicate names\", func(t *", testingT, ") {")
		f.P("fx.maybeSkip(t)")
		methodBatchGet{
			resource: r.resource,
			method:   batchGetMethod,
			parent:   "parent",
			names:    []string{"created00.Name", "created00.Name"},
		}.Generate(f, "response", "err", ":=")
		f.P(assertNilError, "(t, err)")
		f.P(assertDeepEqual, "(")
		f.P("t,")
		f.P("[]*", r.message.GoIdent, "{")
		f.P("created00,")
		f.P("created00,")
		f.P("},")
		f.P("response.", responseResources, ",")
		f.P(protocmpTransform, "(),")
		f.P(")")
		f.P("})")
		f.P()
		// TODO: add test for supplying wildcard as name

		f.P("_ = ", codesNotFound)
	})
}
