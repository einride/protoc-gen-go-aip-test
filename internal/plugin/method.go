package plugin

import (
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/protobuf/compiler/protogen"
)

type methodCreate struct {
	resource *aipreflect.ResourceDescriptor
	method   *protogen.Method

	parent         string
	message        string
	userSettableID string
}

func (m *methodCreate) Generate(f *protogen.GeneratedFile, response string, err string, assign string) {
	f.P(response, ", ", err, " ", assign, " fx.service.", m.method.GoName, "(fx.ctx, &", m.method.Input.GoIdent, "{")
	if hasParent(m.resource) {
		f.P("Parent: ", m.parent, ",")
	}

	switch {
	case m.message != "":
		f.P(m.resource.Singular.UpperCamelCase(), ": ", m.message, ",")
	case !hasParent(m.resource):
		f.P(m.resource.Singular.UpperCamelCase(), ": fx.Create(),")
	default:
		f.P(m.resource.Singular.UpperCamelCase(), ": fx.Create(", m.parent, "),")
	}

	if hasUserSettableID(m.resource, m.method.Desc) && m.userSettableID != "" {
		f.P(m.resource.Singular.UpperCamelCase(), "Id: ", m.userSettableID, ",")
	}
	f.P("})")
}
