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
	s.generateConfigSupplierInterface(f)
	s.generateMainRunMethod(f)
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
	serviceFx := serviceTestSuiteName(s.service.Desc)
	for _, resource := range s.resources {
		resourceFx := resourceTestSuiteConfigName(s.service.Desc, resource)
		f.P("func (fx ", serviceFx, ") Test", resourceType(resource), "(ctx ", context, ", options ", resourceFx, ") {")
		f.P("fx.T.Run(", strconv.Quote(resourceType(resource)), ", func(t *", testingT, ") {")
		f.P("options.ctx = ctx")
		f.P("options.Context = func() ", context, "{ return ctx }")
		f.P("options.service = fx.Server")

		f.P("options.test(t)")
		f.P("})")
		f.P("}")
		f.P()
	}
}

func (s *serviceGenerator) generateConfigSupplierInterface(f *protogen.GeneratedFile) {
	t := f.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "T",
		GoImportPath: "testing",
	})
	f.P("type ", serviceTestConfigSupplierName(s.service.Desc), " interface {")
	for _, resource := range s.resources {
		resourceFx := resourceTestSuiteConfigName(s.service.Desc, resource)
		f.P("Get", resourceType(resource), "TestConfig(t *", t, ") *", resourceFx, "")
	}
	f.P("}")
}

func (s *serviceGenerator) generateMainRunMethod(f *protogen.GeneratedFile) {
	t := f.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "T",
		GoImportPath: "testing",
	})
	f.P("func Test", string(s.service.Desc.Name()), "(")
	f.P("t *", t, ",")
	f.P("s ", serviceTestConfigSupplierName(s.service.Desc), ",")
	f.P(") {")
	for _, resource := range s.resources {
		f.P("test", resourceType(resource), "(t, s)")
	}
	f.P("}")
	f.P()
}

func (s *serviceGenerator) generateTestFunctions(f *protogen.GeneratedFile) {
	t := f.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "T",
		GoImportPath: "testing",
	})
	for _, resource := range s.resources {
		resourceFx := resourceTestSuiteConfigName(s.service.Desc, resource)
		_ = resourceFx
		f.P("func test", resourceType(resource), "(")
		f.P("t *", t, ",")
		f.P("s ", serviceTestConfigSupplierName(s.service.Desc), ",")
		f.P(") {")
		f.P("t.Run(", strconv.Quote(resourceType(resource)), ", func(t *", t, ") {")
		f.P("config := s.Get", resourceType(resource), "TestConfig(t)")
		f.P("if (config == nil) {")
		f.P("t.Skip(\"Method ", "Get", resourceType(resource), "TestConfig not implemented\")")
		f.P("}")
		f.P("config.test(t)")
		f.P("})")
		f.P("}")
		f.P()
	}
}
