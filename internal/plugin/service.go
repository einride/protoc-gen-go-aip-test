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
	return nil
}

func (s *serviceGenerator) generateFixture(f *protogen.GeneratedFile) {
	f.P("type ", s.service.GoName, " struct {")
	f.P("}")
}
