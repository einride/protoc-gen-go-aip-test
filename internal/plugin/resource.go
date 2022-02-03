package plugin

import (
	"strconv"

	"github.com/einride/protoc-gen-go-aip-test/internal/aiptest"
	"github.com/einride/protoc-gen-go-aip-test/internal/ident"
	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
	"github.com/einride/protoc-gen-go-aip-test/internal/util"
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
)

type resourceGenerator struct {
	service  *protogen.Service
	resource *annotations.ResourceDescriptor
	message  *protogen.Message
}

func (r *resourceGenerator) Generate(f *protogen.GeneratedFile) error {
	r.generateFixture(f)
	r.generateTestMethod(f)
	if err := r.generateTestCases(f); err != nil {
		return err
	}
	r.generateParentMethods(f)
	r.generateSkip(f)
	return nil
}

func (r *resourceGenerator) generateFixture(f *protogen.GeneratedFile) {
	context := f.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "Context",
		GoImportPath: "context",
	})
	service := f.QualifiedGoIdent(protogen.GoIdent{
		GoName:       r.service.GoName + "Server",
		GoImportPath: r.service.Methods[0].Input.GoIdent.GoImportPath,
	})

	f.P("type ", resourceTestSuiteConfigName(r.resource), " struct {")
	f.P("ctx ", context)
	f.P("service ", service)
	f.P("currParent int")
	f.P()

	if util.HasParent(r.resource) {
		f.P("// The parents to use when creating resources.")
		f.P("// At least one parent needs to be set. Depending on methods available on the resource,")
		f.P("// more may be required. If insufficient number of parents are")
		f.P("// provided the test will fail.")
		f.P("Parents []string")
	}
	_, hasCreate := util.StandardMethod(r.service, r.resource, aipreflect.MethodTypeCreate)
	if hasCreate {
		f.P("// Create should return a resource which is valid to create, i.e.")
		f.P("// all required fields set.")
		if util.HasParent(r.resource) {
			f.P("Create func(parent string) *", r.message.GoIdent)
		} else {
			f.P("Create func() *", r.message.GoIdent)
		}
	}
	_, hasUpdate := util.StandardMethod(r.service, r.resource, aipreflect.MethodTypeUpdate)
	if hasUpdate {
		f.P("// Update should return a resource which is valid to update, i.e.")
		f.P("// all required fields set.")
		if util.HasParent(r.resource) {
			f.P("Update func(parent string) *", r.message.GoIdent)
		} else {
			f.P("Update func() *", r.message.GoIdent)
		}
	}

	f.P("// Patterns of tests to skip.")
	f.P("// For example if a service has a Get method:")
	f.P("// Skip: [\"Get\"] will skip all tests for Get.")
	f.P("// Skip: [\"Get/persisted\"] will only skip the subtest called \"persisted\" of Get.")
	f.P("Skip []string")
	f.P("}")
	f.P()
}

func (r *resourceGenerator) generateTestMethod(f *protogen.GeneratedFile) {
	testingT := f.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "T",
		GoImportPath: "testing",
	})

	f.P("func (fx *", resourceTestSuiteConfigName(r.resource), ") test(t *", testingT, ") {")
	scope := suite.Scope{
		Service:  r.service,
		Resource: r.resource,
		Message:  r.message,
	}
	for _, s := range aiptest.Suites {
		if s.Enabled(scope) {
			f.P("t.Run(", strconv.Quote(s.Name), ", fx.test", s.Name, ")")
		}
	}
	f.P("}")
	f.P()
}

func (r *resourceGenerator) generateTestCases(f *protogen.GeneratedFile) error {
	testingT := f.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "T",
		GoImportPath: "testing",
	})
	scope := suite.Scope{
		Service:  r.service,
		Resource: r.resource,
		Message:  r.message,
	}
	for _, s := range aiptest.Suites {
		if !s.Enabled(scope) {
			continue
		}
		f.P("func (fx *", resourceTestSuiteConfigName(r.resource), ") test", s.Name, "(t *", testingT, ") {")
		f.P(ident.FixtureMaybeSkip, "(t)")
		for _, t := range s.Tests {
			if !t.Enabled(scope) {
				continue
			}
			if err := r.generateTestCase(f, t, scope); err != nil {
				return err
			}
			f.P()
		}
		for _, group := range s.TestGroups {
			if !group.Enabled(scope) {
				continue
			}
			if err := group.GenerateBefore(f, scope); err != nil {
				return err
			}
			for _, t := range group.Tests {
				if !t.Enabled(scope) {
					continue
				}
				if err := r.generateTestCase(f, t, scope); err != nil {
					return err
				}
				f.P()
			}
		}
		f.P("}")
		f.P()
	}
	return nil
}

func (r *resourceGenerator) generateTestCase(f *protogen.GeneratedFile, test suite.Test, scope suite.Scope) error {
	testingT := f.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "T",
		GoImportPath: "testing",
	})
	for _, line := range test.Doc {
		f.P("// ", line)
	}
	f.P("t.Run(", strconv.Quote(test.Name), ", func(t *", testingT, ") {")
	f.P(ident.FixtureMaybeSkip, "(t)")
	if err := test.Generate(f, scope); err != nil {
		return err
	}
	f.P("})")
	return nil
}

func (r *resourceGenerator) generateSkip(f *protogen.GeneratedFile) {
	testingT := f.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "T",
		GoImportPath: "testing",
	})
	stringsContains := f.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "Contains",
		GoImportPath: "strings",
	})
	f.P("func (fx *", resourceTestSuiteConfigName(r.resource), ") maybeSkip(t *", testingT, ") {")
	f.P("for _, skip := range fx.Skip {")
	f.P("if ", stringsContains, "(t.Name(), skip) {")
	f.P("t.Skip(\"skipped because of .Skip\")")
	f.P("}")
	f.P("}")
	f.P("}")
	f.P()
}

func (r *resourceGenerator) generateParentMethods(f *protogen.GeneratedFile) {
	if !util.HasParent(r.resource) {
		return
	}
	testingT := f.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "T",
		GoImportPath: "testing",
	})
	f.P("func (fx *", resourceTestSuiteConfigName(r.resource), ") nextParent(t *", testingT, ", pristine bool) string {")
	f.P("if pristine {")
	f.P("fx.currParent++")
	f.P("}")
	f.P("if fx.currParent >= len(fx.Parents) {")
	f.P("t.Fatal(\"need at least\", fx.currParent + 1,  \"parents\")")
	f.P("}")
	f.P("return fx.Parents[fx.currParent]")
	f.P("}")
	f.P()
	f.P("func (fx *", resourceTestSuiteConfigName(r.resource), ") peekNextParent(t *", testingT, ") string {")
	f.P("next := fx.currParent + 1")
	f.P("if next >= len(fx.Parents) {")
	f.P("t.Fatal(\"need at least\", next +1,  \"parents\")")
	f.P("}")
	f.P("return fx.Parents[next]")
	f.P("}")
	f.P()
}
