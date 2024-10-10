package plugin

import (
	"strconv"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
)

type serviceGenerator struct {
	service   *protogen.Service
	resources []*annotations.ResourceDescriptor
	messages  []*protogen.Message
}

func (s *serviceGenerator) Generate(f *protogen.GeneratedFile) error {
	s.generateConfigProvider(f)
	s.generateMainTestFunction(f)
	s.generateTestFunctions(f)
	s.generateFixture(f)
	s.generateTestMethods(f)
	for i, resource := range s.resources {
		message := s.messages[i]
		generator := resourceGenerator{
			service:  s.service,
			resource: resource,
			message:  message,
		}
		if err := generator.Generate(f); err != nil {
			return err
		}
	}
	return nil
}

func (s *serviceGenerator) generateConfigProvider(f *protogen.GeneratedFile) {
	t := f.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "T",
		GoImportPath: "testing",
	})
	name := serviceTestConfigProviderName(s.service.Desc)
	f.P("// ", name, " is the interface to implement to decide which resources")
	f.P("// that should be tested and how it's configured.")
	f.P("type ", name, " interface {")
	for _, resource := range s.resources {
		resourceFx := serviceResourceName(s.service.Desc, resource)
		f.P("// ", resourceFx, " should return a config, or nil, which means that the tests will be skipped.")
		f.P(resourceFx, "(t *", t, ") *", resourceTestSuiteConfigName(s.service.Desc, resource), "")
	}
	f.P("}")
	f.P()
}

func (s *serviceGenerator) generateMainTestFunction(f *protogen.GeneratedFile) {
	t := f.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "T",
		GoImportPath: "testing",
	})
	funcName := "Test" + string(s.service.Desc.Name())
	f.P("// ", funcName, " is the main entrypoint for starting the AIP tests.")
	f.P("func ", funcName, "(t *", t, ",s ", serviceTestConfigProviderName(s.service.Desc), ") {")
	for _, resource := range s.resources {
		name := serviceResourceName(s.service.Desc, resource)
		f.P("test", name, "(t, s)")
	}
	f.P("}")
	f.P()
}

func (s *serviceGenerator) generateTestFunctions(f *protogen.GeneratedFile) {
	t := f.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "T",
		GoImportPath: "testing",
	})
	context := f.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "Context",
		GoImportPath: "context",
	})
	background := f.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "Background",
		GoImportPath: "context",
	})
	for _, resource := range s.resources {
		name := serviceResourceName(s.service.Desc, resource)
		f.P("func test", name, "(t *", t, ",s ", serviceTestConfigProviderName(s.service.Desc), ") {")
		f.P("t.Run(", strconv.Quote(resourceType(resource)), ", func(t *", t, ") {")
		f.P("config := s.", serviceResourceName(s.service.Desc, resource), "(t)")
		f.P("if (config == nil) {")
		f.P("t.Skip(\"Method ", name, " not implemented\")")
		f.P("}")
		f.P("if (config.Service == nil) {")
		f.P("t.Skip(\"Method ", name, ".Service() not implemented\")")
		f.P("}")
		f.P("if (config.Context == nil) {")
		f.P("config.Context = func() ", context, " { return ", background, "() }")
		f.P("}")
		f.P("config.test(t)")
		f.P("})")
		f.P("}")
		f.P()
	}
}

func (s *serviceGenerator) generateFixture(f *protogen.GeneratedFile) {
	testingT := f.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "T",
		GoImportPath: "testing",
	})

	service := f.QualifiedGoIdent(protogen.GoIdent{
		GoName:       s.service.GoName + "Server",
		GoImportPath: s.service.Methods[0].Input.GoIdent.GoImportPath,
	})

	f.P("type ", serviceTestSuiteName(s.service.Desc), " struct {")
	f.P("T *", testingT)
	f.P("// Server to test.")
	f.P("Server  ", service)
	f.P()

	f.P("}")
	f.P()
}

func (s *serviceGenerator) generateTestMethods(f *protogen.GeneratedFile) {
	context := f.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "Context",
		GoImportPath: "context",
	})
	testingT := f.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "T",
		GoImportPath: "testing",
	})
	service := f.QualifiedGoIdent(protogen.GoIdent{
		GoName:       s.service.GoName + "Server",
		GoImportPath: s.service.Methods[0].Input.GoIdent.GoImportPath,
	})
	serviceFx := serviceTestSuiteName(s.service.Desc)
	for _, resource := range s.resources {
		resourceFx := resourceTestSuiteConfigName(s.service.Desc, resource)
		f.P("func (fx ", serviceFx, ") Test", resourceType(resource), "(ctx ", context, ", options ", resourceFx, ") {")
		f.P("fx.T.Run(", strconv.Quote(resourceType(resource)), ", func(t *", testingT, ") {")
		f.P("options.Context = func() ", context, " { return ctx }")
		f.P("options.Service = func() ", service, " { return fx.Server", "}")
		f.P("options.test(t)")
		f.P("})")
		f.P("}")
		f.P()
	}
}
