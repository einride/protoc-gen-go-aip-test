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

	f.P("type ", s.service.GoName, " struct {")
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
	serviceFx := s.service.GoName
	for _, resource := range s.resources {
		resourceFx := resourceType(resource)
		f.P("func (fx *", serviceFx, ") Test", resourceFx, "(ctx ", context, ", options ", resourceFx, ") {")
		f.P("fx.T.Run(", strconv.Quote(resourceType(resource)), ", func(t *", testingT, ") {")
		f.P("options.ctx = ctx")
		f.P("options.service = fx.Server")
		f.P("options.test(t)")
		f.P("})")
		f.P("}")
		f.P()
	}
}
