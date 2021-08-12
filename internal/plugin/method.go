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

func (m methodCreate) Generate(f *protogen.GeneratedFile, response string, err string, assign string) {
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

type methodGet struct {
	resource *annotations.ResourceDescriptor
	method   *protogen.Method

	name string
}

func (m methodGet) Generate(f *protogen.GeneratedFile, response string, err string, assign string) {
	f.P(response, ", ", err, " ", assign, " fx.service.", m.method.GoName, "(fx.ctx, &", m.method.Input.GoIdent, "{")
	f.P("Name: ", m.name, ",")
	f.P("})")
}

type methodBatchGet struct {
	resource *annotations.ResourceDescriptor
	method   *protogen.Method

	parent string
	names  []string
}

func (m methodBatchGet) Generate(f *protogen.GeneratedFile, response string, err string, assign string) {
	f.P(response, ", ", err, " ", assign, " fx.service.", m.method.GoName, "(fx.ctx, &", m.method.Input.GoIdent, "{")
	if hasParent(m.resource) {
		f.P("Parent: ", m.parent, ",")
	}
	f.P("Names: []string{")
	for _, name := range m.names {
		f.P(name, ",")
	}
	f.P("},")
	f.P("})")
}

type methodUpdate struct {
	resource *annotations.ResourceDescriptor
	method   *protogen.Method

	// set either parent + name, or msg
	name       string
	parent     string
	msg        string
	updateMask []string
}

func (m methodUpdate) Generate(f *protogen.GeneratedFile, response string, err string, assign string) {
	upper := aipreflect.GrammaticalName(m.resource.GetSingular()).UpperCamelCase()

	if m.msg == "" {
		if hasParent(m.resource) {
			f.P("msg := fx.Update(", m.parent, ")")
		} else {
			f.P("msg := fx.Update()")
		}
		f.P("msg.Name = ", m.name)
	}
	f.P(response, ", ", err, " ", assign, " fx.service.", m.method.GoName, "(fx.ctx, &", m.method.Input.GoIdent, "{")
	if m.msg != "" {
		f.P(upper, ":", m.msg, ",")
	} else {
		f.P(upper, ": msg,")
	}
	if hasUpdateMask(m.method.Desc) && len(m.updateMask) > 0 {
		fieldmaskpbFieldMask := f.QualifiedGoIdent(protogen.GoIdent{
			GoName:       "FieldMask",
			GoImportPath: "google.golang.org/protobuf/types/known/fieldmaskpb",
		})
		f.P("UpdateMask: &", fieldmaskpbFieldMask, "{")
		f.P("Paths: []string{")
		for _, path := range m.updateMask {
			f.P(path, ",")
		}
		f.P("},")
		f.P("},")
	}
	f.P("})")
}
