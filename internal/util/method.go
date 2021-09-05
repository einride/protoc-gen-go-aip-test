package util

import (
	"github.com/stoewer/go-strcase"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
)

type MethodCreate struct {
	Resource *annotations.ResourceDescriptor
	Method   *protogen.Method

	Parent         string
	Message        string
	UserSettableID string
}

func (m MethodCreate) Generate(f *protogen.GeneratedFile, response string, err string, assign string) {
	f.P(response, ", ", err, " ", assign, " fx.service.", m.Method.GoName, "(fx.ctx, &", m.Method.Input.GoIdent, "{")
	if HasParent(m.Resource) {
		f.P("Parent: ", m.Parent, ",")
	}

	upper := strcase.UpperCamelCase(string(FindResourceField(
		m.Method.Input.Desc,
		m.Resource,
	).Name()))

	switch {
	case m.Message != "":
		f.P(upper, ": ", m.Message, ",")
	case !HasParent(m.Resource):
		f.P(upper, ": fx.Create(),")
	default:
		f.P(upper, ": fx.Create(", m.Parent, "),")
	}

	if m.UserSettableID != "" && HasUserSettableIDField(m.Resource, m.Method.Input.Desc) {
		f.P(upper, "Id: ", m.UserSettableID, ",")
	}
	f.P("})")
}

type MethodGet struct {
	Resource *annotations.ResourceDescriptor
	Method   *protogen.Method

	Name string
}

func (m MethodGet) Generate(f *protogen.GeneratedFile, response string, err string, assign string) {
	f.P(response, ", ", err, " ", assign, " fx.service.", m.Method.GoName, "(fx.ctx, &", m.Method.Input.GoIdent, "{")
	f.P("Name: ", m.Name, ",")
	f.P("})")
}

type MethodBatchGet struct {
	Resource *annotations.ResourceDescriptor
	Method   *protogen.Method

	Parent string
	Names  []string
}

func (m MethodBatchGet) Generate(f *protogen.GeneratedFile, response string, err string, assign string) {
	f.P(response, ", ", err, " ", assign, " fx.service.", m.Method.GoName, "(fx.ctx, &", m.Method.Input.GoIdent, "{")
	if HasParent(m.Resource) {
		f.P("Parent: ", m.Parent, ",")
	}
	f.P("Names: []string{")
	for _, name := range m.Names {
		f.P(name, ",")
	}
	f.P("},")
	f.P("})")
}

type MethodUpdate struct {
	Resource *annotations.ResourceDescriptor
	Method   *protogen.Method

	// set either Parent + Name, or Msg
	Name       string
	Parent     string
	Msg        string
	UpdateMask []string
}

func (m MethodUpdate) Generate(f *protogen.GeneratedFile, response string, err string, assign string) {
	upper := strcase.UpperCamelCase(string(FindResourceField(
		m.Method.Input.Desc,
		m.Resource,
	).Name()))

	if m.Msg == "" {
		if HasParent(m.Resource) {
			f.P("msg := fx.Update(", m.Parent, ")")
		} else {
			f.P("msg := fx.Update()")
		}
		f.P("msg.Name = ", m.Name)
	}
	f.P(response, ", ", err, " ", assign, " fx.service.", m.Method.GoName, "(fx.ctx, &", m.Method.Input.GoIdent, "{")
	if m.Msg != "" {
		f.P(upper, ":", m.Msg, ",")
	} else {
		f.P(upper, ": msg,")
	}
	if HasUpdateMask(m.Method.Desc) && len(m.UpdateMask) > 0 {
		fieldmaskpbFieldMask := f.QualifiedGoIdent(protogen.GoIdent{
			GoName:       "FieldMask",
			GoImportPath: "google.golang.org/protobuf/types/known/fieldmaskpb",
		})
		f.P("UpdateMask: &", fieldmaskpbFieldMask, "{")
		f.P("Paths: []string{")
		for _, path := range m.UpdateMask {
			f.P(path, ",")
		}
		f.P("},")
		f.P("},")
	}
	f.P("})")
}

type MethodList struct {
	Resource *annotations.ResourceDescriptor
	Method   *protogen.Method

	Parent    string
	PageSize  string
	PageToken string
}

func (m MethodList) Generate(f *protogen.GeneratedFile, response string, err string, assign string) {
	f.P(response, ", ", err, " ", assign, " fx.service.", m.Method.GoName, "(fx.ctx, &", m.Method.Input.GoIdent, "{")
	if HasParent(m.Resource) {
		f.P("Parent: ", m.Parent, ",")
	}
	if m.PageSize != "" {
		f.P("PageSize: ", m.PageSize, ",")
	}
	if m.PageToken != "" {
		f.P("PageToken: ", m.PageToken, ",")
	}
	f.P("})")
}

type MethodSearch struct {
	Resource *annotations.ResourceDescriptor
	Method   *protogen.Method

	Parent    string
	PageSize  string
	PageToken string
}

func (m MethodSearch) Generate(f *protogen.GeneratedFile, response string, err string, assign string) {
	f.P(response, ", ", err, " ", assign, " fx.service.", m.Method.GoName, "(fx.ctx, &", m.Method.Input.GoIdent, "{")
	if HasParent(m.Resource) {
		f.P("Parent: ", m.Parent, ",")
	}
	if m.PageSize != "" {
		f.P("PageSize: ", m.PageSize, ",")
	}
	if m.PageToken != "" {
		f.P("PageToken: ", m.PageToken, ",")
	}
	f.P("})")
}

type MethodDelete struct {
	Resource *annotations.ResourceDescriptor
	Method   *protogen.Method

	Name string
}

func (m MethodDelete) Generate(f *protogen.GeneratedFile, response string, err string, assign string) {
	f.P(response, ", ", err, " ", assign, " fx.service.", m.Method.GoName, "(fx.ctx, &", m.Method.Input.GoIdent, "{")
	f.P("Name: ", m.Name, ",")
	f.P("})")
}
