package plugin

import (
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
)

type methodCreate struct {
	resource *annotations.ResourceDescriptor
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

	upper := aipreflect.GrammaticalName(m.resource.GetSingular()).UpperCamelCase()
	switch {
	case m.message != "":
		f.P(upper, ": ", m.message, ",")
	case !hasParent(m.resource):
		f.P(upper, ": fx.Create(),")
	default:
		f.P(upper, ": fx.Create(", m.parent, "),")
	}

	if m.userSettableID != "" && hasUserSettableIDField(m.resource, m.method.Input.Desc) {
		f.P(upper, "Id: ", m.userSettableID, ",")
	}
	f.P("})")
}
