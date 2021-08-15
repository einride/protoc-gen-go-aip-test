package plugin

import (
	"strconv"

	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/protobuf/compiler/protogen"
)

func (r *resourceGenerator) searchTestCase() testCase {
	searchMethod, ok := r.standardMethod(aipreflect.MethodTypeSearch)
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

	deleteMethod, hasDelete := r.standardMethod(aipreflect.MethodTypeDelete)

	responseResources := aipreflect.GrammaticalName(r.resource.GetPlural()).UpperCamelCase()

	return newTestCase("Search", func(f *protogen.GeneratedFile) {
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
		cmpoptsSortSlices := f.QualifiedGoIdent(protogen.GoIdent{
			GoName:       "SortSlices",
			GoImportPath: "github.com/google/go-cmp/cmp/cmpopts",
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

		f.P("// Standard methods: Search")
		f.P("// https://google.aip.dev/132")

		if hasParent(r.resource) {
			f.P("parent01 := fx.nextParent(t, false)")
			f.P("parent02 := fx.nextParent(t, true)")
			f.P()
		} else {
		}

		// create 15 under each parent
		f.P("const n = 15")
		f.P()
		if hasParent(r.resource) {
			f.P("parent01msgs := make([]*", r.message.GoIdent, ", n)")
			f.P("for i := 0; i < n; i++ {")
			methodCreate{
				resource: r.resource,
				method:   createMethod,
				parent:   "parent01",
			}.Generate(f, "msg", "err", ":=")
			f.P(assertNilError, "(t, err)")
			f.P("parent01msgs[i] = msg")
			f.P("}")
			f.P()
		}
		f.P("parent02msgs := make([]*", r.message.GoIdent, ", n)")
		f.P("for i := 0; i < n; i++ {")
		methodCreate{
			resource: r.resource,
			method:   createMethod,
			parent:   "parent02",
		}.Generate(f, "msg", "err", ":=")
		f.P(assertNilError, "(t, err)")
		f.P("parent02msgs[i] = msg")
		f.P("}")

		if hasParent(r.resource) {
			f.P()
			f.P("// Method should fail with InvalidArgument is provided parent is not valid.")
			f.P("t.Run(\"invalid parent\", func(t *", testingT, ") {")
			f.P("fx.maybeSkip(t)")
			methodSearch{
				resource: r.resource,
				method:   searchMethod,
				parent:   strconv.Quote("invalid parent"),
			}.Generate(f, "_", "err", ":=")
			f.P(assertEqual, "(t, ", codesInvalidArgument, ",", statusCode, "(err), err)")
			f.P("})")
		}

		f.P()
		f.P("// Method should fail with InvalidArgument is provided page token is not valid.")
		f.P("t.Run(\"invalid page token\", func(t *", testingT, ") {")
		f.P("fx.maybeSkip(t)")
		methodSearch{
			resource:  r.resource,
			method:    searchMethod,
			parent:    "parent01",
			pageToken: strconv.Quote("invalid page token"),
		}.Generate(f, "_", "err", ":=")
		f.P(assertEqual, "(t, ", codesInvalidArgument, ",", statusCode, "(err), err)")
		f.P("})")

		f.P()
		f.P("// Method should fail with InvalidArgument is provided page size is negative.")
		f.P("t.Run(\"negative page size\", func(t *", testingT, ") {")
		f.P("fx.maybeSkip(t)")
		methodSearch{
			resource: r.resource,
			method:   searchMethod,
			parent:   "parent01",
			pageSize: "-10",
		}.Generate(f, "_", "err", ":=")
		f.P(assertEqual, "(t, ", codesInvalidArgument, ",", statusCode, "(err), err)")
		f.P("})")

		if hasParent(r.resource) {
			f.P()
			f.P("// If parent is provided the method must only return resources")
			f.P("// under that parent.")
			f.P("t.Run(\"isolation\", func(t *", testingT, ") {")
			f.P("fx.maybeSkip(t)")
			methodSearch{
				resource: r.resource,
				method:   searchMethod,
				parent:   "parent02",
				pageSize: "999",
			}.Generate(f, "response", "err", ":=")
			f.P(assertNilError, "(t, err)")
			f.P(assertDeepEqual, "(")
			f.P("t,")
			f.P("parent02msgs,")
			f.P("response.", responseResources, ",")
			f.P(cmpoptsSortSlices, "(func(a,b *", r.message.GoIdent, ") bool {")
			f.P("return a.Name < b.Name")
			f.P("}),")
			f.P(protocmpTransform, "(),")
			f.P(")")
			f.P("})")
		}

		if hasParent(r.resource) {
			f.P()
			f.P("t.Run(\"pagination\", func(t *", testingT, ") {")
			f.P("fx.maybeSkip(t)")

			f.P()
			f.P("// If there are no more resources, next_page_token should be unset.")
			f.P("t.Run(\"next page token\", func(t *", testingT, ") {")
			f.P("fx.maybeSkip(t)")
			methodSearch{
				resource: r.resource,
				method:   searchMethod,
				parent:   "parent02",
				pageSize: "999",
			}.Generate(f, "response", "err", ":=")
			f.P(assertNilError, "(t, err)")
			f.P("assert.Equal(t, \"\", response.NextPageToken)")
			f.P("})")
			f.P()

			f.P("// Searching resource one by one should eventually return all resources created.")
			f.P("t.Run(\"one by one\", func(t *", testingT, ") {")
			f.P("fx.maybeSkip(t)")
			f.P("msgs := make([]*", r.message.GoIdent, ", 0, n)")
			f.P("var nextPageToken string")
			f.P("for {")
			methodSearch{
				resource: r.resource,
				method:   searchMethod,
				parent:   "parent02",
				pageSize: "1",
			}.Generate(f, "response", "err", ":=")
			f.P(assertNilError, "(t, err)")
			f.P(assertEqual, "(t, 1, len(response.", responseResources, "))")
			f.P("msgs = append(msgs, response.", responseResources, "...)")
			f.P("nextPageToken = response.NextPageToken")
			f.P("if nextPageToken == \"\" {")
			f.P("break")
			f.P("}")
			f.P("}")
			f.P(assertDeepEqual, "(")
			f.P("t,")
			f.P("parent02msgs,")
			f.P("msgs,")
			f.P(cmpoptsSortSlices, "(func(a,b *", r.message.GoIdent, ") bool {")
			f.P("return a.Name < b.Name")
			f.P("}),")
			f.P(protocmpTransform, "(),")
			f.P(")")
			f.P("})")
			f.P("})")
			f.P()
		}

		if hasParent(r.resource) && hasDelete {
			f.P()
			f.P("// Method should not return deleted resources.")
			f.P("t.Run(\"deleted\", func(t *", testingT, ") {")
			f.P("fx.maybeSkip(t)")
			f.P("const nDelete = 5")
			f.P("for i := 0; i < nDelete; i++ {")
			methodDelete{
				method:   deleteMethod,
				resource: r.resource,
				name:     "parent02msgs[i].Name",
			}.Generate(f, "_", "err", ":=")
			f.P(assertNilError, "(t, err)")
			f.P("}")
			methodSearch{
				resource: r.resource,
				method:   searchMethod,
				parent:   "parent02",
				pageSize: "9999",
			}.Generate(f, "response", "err", ":=")
			f.P(assertNilError, "(t, err)")
			f.P(assertDeepEqual, "(")
			f.P("t,")
			f.P("parent02msgs[nDelete:],")
			f.P("response.", responseResources, ",")
			f.P(cmpoptsSortSlices, "(func(a,b *", r.message.GoIdent, ") bool {")
			f.P("return a.Name < b.Name")
			f.P("}),")
			f.P(protocmpTransform, "(),")
			f.P(")")
			f.P("})")

		}

		f.P("_ = ", codesNotFound)
		f.P("_ = ", protocmpTransform)
		f.P("_ = ", cmpoptsSortSlices)
	})
}
