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
	r.generateCreate(f)
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

	f.P("type ", resourceTestSuiteConfigName(r.service.Desc, r.resource), " struct {")
	f.P("currParent int")
	f.P()

	f.P("// Service should return the service that should be tested.")
	f.P("// The service will be used for several tests.")
	f.P("Service", " func() ", service)
	f.P("// Context should return a new context.")
	f.P("// The context will be used for several tests.")
	f.P("Context", " func() ", context)
	if util.HasParent(r.resource) {
		f.P("// The parents to use when creating resources.")
		f.P("// At least one parent needs to be set. Depending on methods available on the resource,")
		f.P("// more may be required. If insufficient number of parents are")
		f.P("// provided the test will fail.")
		f.P("Parents []string")
	}
	createMethod, hasCreate := util.StandardMethod(r.service, r.resource, aipreflect.MethodTypeCreate)
	if hasCreate {
		f.P("// Create should return a resource which is valid to create, i.e.")
		f.P("// all required fields set.")
		if util.HasParent(r.resource) {
			f.P("Create func(parent string) *", r.message.GoIdent)
		} else {
			f.P("Create func() *", r.message.GoIdent)
		}

		if util.HasUserSettableIDField(r.resource, createMethod.Input.Desc) {
			f.P("// IDGenerator should return a valid and unique ID to use in the Create call.")
			f.P("// If non-nil, this function will be called to set the ID on all Create calls.")
			f.P("// If the ID field is required, tests will fail if this is nil.")
			f.P("IDGenerator func() string")
		}
	} else {
		f.P("// CreateResource should create a ", r.message.Desc.Name(), " and return it.")
		f.P("// If the field is not set, some tests will be skipped.")
		f.P("//")
		f.P("// This method is generated because service does not expose a Create")
		f.P("// method (or it does not comply with AIP).")
		if util.HasParent(r.resource) {
			f.P("CreateResource func(ctx ", context, ", parent string) (*", r.message.GoIdent, ", error)")
		} else {
			f.P("CreateResource func(ctx ", context, ") (*", r.message.GoIdent, ", error)")
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

	f.P("func (fx *", resourceTestSuiteConfigName(r.service.Desc, r.resource), ") test(t *", testingT, ") {")
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
		f.P("func (fx *", resourceTestSuiteConfigName(r.service.Desc, r.resource), ") test", s.Name, "(t *", testingT, ") {")
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
			// Create a new block for each group to not conflict with each other.
			f.P("{")
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
			f.P("}")
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
	f.P("func (fx *", resourceTestSuiteConfigName(r.service.Desc, r.resource), ") maybeSkip(t *", testingT, ") {")
	f.P("for _, skip := range fx.Skip {")
	f.P("if ", stringsContains, "(t.Name(), skip) {")
	f.P("t.Skip(\"skipped because of .Skip\")")
	f.P("}")
	f.P("}")
	f.P("}")
	f.P()
}

func (r *resourceGenerator) generateCreate(f *protogen.GeneratedFile) {
	testingT := f.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "T",
		GoImportPath: "testing",
	})
	fixtureName := resourceTestSuiteConfigName(r.service.Desc, r.resource)
	var parentFuncArg string
	var parentCallArg string
	if util.HasParent(r.resource) {
		parentFuncArg = ", parent string"
		parentCallArg = ", parent"
	}
	createMethod, hasCreate := util.StandardMethod(r.service, r.resource, aipreflect.MethodTypeCreate)
	isLROCreate := hasCreate && util.ReturnsLRO(createMethod.Desc)

	f.P("func (fx *", fixtureName, ") create(t *", testingT, parentFuncArg, ") *", r.message.GoIdent, "{")
	f.P("t.Helper()")
	switch {
	case hasCreate && isLROCreate:
		f.P("t.Skip(\"Long running create method not supported\")")
		f.P("return nil")
	case hasCreate:
		util.MethodCreate{
			Resource: r.resource,
			Method:   createMethod,
			Parent:   "parent",
		}.Generate(f, "created", "err", ":=")
		f.P(ident.AssertNilError, "(t, err)")
		f.P("return created")
	default:
		f.P("if fx.CreateResource == nil {")
		f.P("t.Skip(\"Test skipped because CreateResource not specified on ", fixtureName, "\")")
		f.P("}")
		f.P("created, err := fx.CreateResource(fx.Context()", parentCallArg, ")")
		f.P(ident.AssertNilError, "(t, err)")
		f.P("return created")
	}
	f.P("}")
}

func (r *resourceGenerator) generateParentMethods(f *protogen.GeneratedFile) {
	if !util.HasParent(r.resource) {
		return
	}
	testingT := f.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "T",
		GoImportPath: "testing",
	})
	f.P("func (fx *", resourceTestSuiteConfigName(
		r.service.Desc,
		r.resource,
	), ") nextParent(t *", testingT, ", pristine bool) string {")
	f.P("if pristine {")
	f.P("fx.currParent++")
	f.P("}")
	f.P("if fx.currParent >= len(fx.Parents) {")
	f.P("t.Fatal(\"need at least\", fx.currParent + 1,  \"parents\")")
	f.P("}")
	f.P("return fx.Parents[fx.currParent]")
	f.P("}")
	f.P()
	f.P("func (fx *", resourceTestSuiteConfigName(
		r.service.Desc,
		r.resource,
	), ") peekNextParent(t *", testingT, ") string {")
	f.P("next := fx.currParent + 1")
	f.P("if next >= len(fx.Parents) {")
	f.P("t.Fatal(\"need at least\", next +1,  \"parents\")")
	f.P("}")
	f.P("return fx.Parents[next]")
	f.P("}")
	f.P()
}
