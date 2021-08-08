package plugin

import (
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type serviceGenerator struct {
	service   *protogen.Service
	resources []*aipreflect.ResourceDescriptor
	messages  []protoreflect.MessageDescriptor
}

func (s *serviceGenerator) Generate(f *protogen.GeneratedFile) error {
	s.generateFixture(f)
	s.generateTestMethods(f)
	for i, resource := range s.resources {
		message := s.messages[i]
		generator := resourceGenerator{
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
	context := f.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "Context",
		GoImportPath: "context",
	})

	service := f.QualifiedGoIdent(protogen.GoIdent{
		GoName:       s.service.GoName + "Server",
		GoImportPath: s.service.Methods[0].Input.GoIdent.GoImportPath,
	})

	f.P("type ", s.service.GoName, " struct {")

	f.P("// Context to use for running tests.")
	f.P("Context ", context)
	f.P()

	f.P("// Service to test.")
	f.P("Service  ", service)
	f.P()

	f.P("}")
	f.P()
}

func (s *serviceGenerator) generateTestMethods(f *protogen.GeneratedFile) {
	testing := f.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "T",
		GoImportPath: "testing",
	})
	serviceFx := s.service.GoName
	for _, resource := range s.resources {
		resourceFx := resource.Type.Type()
		f.P("func (fx *", serviceFx, ") Test", resourceFx, "(t *", testing, ", options ", resourceFx, ") {")
		f.P("options.test(t)")
		f.P("}")
		f.P()
	}
}
