package plugin

import (
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type resourceGenerator struct {
	resource *aipreflect.ResourceDescriptor
	message  protoreflect.MessageDescriptor
}

func (r *resourceGenerator) Generate(f *protogen.GeneratedFile) error {
	r.generateFixture(f)
	r.generateTestMethod(f)
	return nil
}

func (r *resourceGenerator) generateFixture(f *protogen.GeneratedFile) {
	f.P("type ", r.resource.Type.Type(), " struct {")
	f.P("}")
	f.P()
}

func (r *resourceGenerator) generateTestMethod(f *protogen.GeneratedFile) {
	testing := f.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "T",
		GoImportPath: "testing",
	})

	f.P("func (fx *", r.resource.Type.Type(), ") test(t *", testing, ") {")
	f.P("}")
}
