package util

import (
	"github.com/einride/protoc-gen-go-aip-test/internal/transport"
	"github.com/stoewer/go-strcase"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
)

// beginCall emits the opening of a service method call.
// For gRPC/connect-simple: response, err := fx.Service().Method(fx.Context(), &Request{
// For connect standard:    _resp, err := fx.Service().Method(fx.Context(), connect.NewRequest(&Request{
// Returns the actual response variable name used in the emitted code.
func beginCall(
	f *protogen.GeneratedFile,
	t transport.Transport,
	method *protogen.Method,
	response, err, assign string,
) {
	if t.UsesRequestWrapper() {
		actualResp := response
		if response != "_" {
			actualResp = "_resp"
		}
		f.P(
			actualResp,
			", ",
			err,
			" ",
			assign,
			" fx.Service().",
			method.GoName,
			"(fx.Context(), ",
			f.QualifiedGoIdent(t.NewRequestIdent()),
			"(&",
			method.Input.GoIdent,
			"{",
		)
	} else {
		f.P(
			response,
			", ",
			err,
			" ",
			assign,
			" fx.Service().",
			method.GoName,
			"(fx.Context(), &",
			method.Input.GoIdent,
			"{",
		)
	}
}

// endCall emits the closing of a service method call and, for connect standard mode,
// unwraps the response from connect.Response[T].
func endCall(
	f *protogen.GeneratedFile,
	t transport.Transport,
	method *protogen.Method,
	response string,
) {
	if t.UsesRequestWrapper() {
		f.P("}))")
		if response != "_" {
			f.P("var ", response, " *", method.Output.GoIdent)
			f.P("if _resp != nil { ", response, " = _resp.Msg }")
		}
	} else {
		f.P("})")
	}
}

type MethodCreate struct {
	Resource *annotations.ResourceDescriptor
	Method   *protogen.Method

	Parent         string
	Message        string
	UserSettableID string
}

func (m MethodCreate) Generate(f *protogen.GeneratedFile, t transport.Transport, response, err, assign string) {
	userSetID := m.UserSettableID
	if userSetID == "" && HasUserSettableIDField(m.Resource, m.Method.Input.Desc) {
		userSetID = "userSetID"
		f.P(userSetID + " := \"\"")
		f.P("if fx.IDGenerator != nil {")
		f.P(userSetID + " = fx.IDGenerator()")
		f.P("}")
	}

	beginCall(f, t, m.Method, response, err, assign)
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

	if userSetID != "" && HasUserSettableIDField(m.Resource, m.Method.Input.Desc) {
		f.P(upper, "Id: ", userSetID, ",")
	}

	endCall(f, t, m.Method, response)
}

type MethodGet struct {
	Resource *annotations.ResourceDescriptor
	Method   *protogen.Method

	Name string
}

func (m MethodGet) Generate(f *protogen.GeneratedFile, t transport.Transport, response, err, assign string) {
	beginCall(f, t, m.Method, response, err, assign)
	f.P("Name: ", m.Name, ",")
	endCall(f, t, m.Method, response)
}

type MethodBatchGet struct {
	Resource *annotations.ResourceDescriptor
	Method   *protogen.Method

	Parent string
	Names  []string
}

func (m MethodBatchGet) Generate(f *protogen.GeneratedFile, t transport.Transport, response, err, assign string) {
	beginCall(f, t, m.Method, response, err, assign)
	if HasParent(m.Resource) {
		f.P("Parent: ", m.Parent, ",")
	}
	f.P("Names: []string{")
	for _, name := range m.Names {
		f.P(name, ",")
	}
	f.P("},")
	endCall(f, t, m.Method, response)
}

type MethodUpdate struct {
	Resource *annotations.ResourceDescriptor
	Method   *protogen.Method

	// set either Parent + Name, or Msg.
	Name       string
	Parent     string
	Msg        string
	UpdateMask []string
	Etag       string
	EtagTest   bool
}

func (m MethodUpdate) Generate(f *protogen.GeneratedFile, t transport.Transport, response, err, assign string) {
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
	if m.EtagTest && !HasEtagField(m.Method.Input.Desc) && HasEtagField(m.Method.Output.Desc) {
		// Request object does not have an etag field, but the resource has.
		if m.Etag != "" {
			f.P("msg.Etag = ", m.Etag)
		} else {
			f.P(`msg.Etag = created.Etag // assign etag from the created resource`)
		}
	}
	beginCall(f, t, m.Method, response, err, assign)
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
	switch {
	case HasEtagField(m.Method.Input.Desc) && m.Etag != "":
		f.P("Etag: ", m.Etag, ",")
	case HasRequiredEtagField(m.Method.Input.Desc):
		if m.Msg != "" {
			// Delete request has an required etag field.
			f.P("Etag: ", m.Msg, ".Etag,")
		} else {
			f.P("Etag: msg.Etag,")
		}
	}
	endCall(f, t, m.Method, response)
}

type MethodList struct {
	Resource *annotations.ResourceDescriptor
	Method   *protogen.Method

	Parent    string
	PageSize  string
	PageToken string
}

func (m MethodList) Generate(f *protogen.GeneratedFile, t transport.Transport, response, err, assign string) {
	beginCall(f, t, m.Method, response, err, assign)
	if HasParent(m.Resource) {
		f.P("Parent: ", m.Parent, ",")
	}
	if m.PageSize != "" {
		f.P("PageSize: ", m.PageSize, ",")
	}
	if m.PageToken != "" {
		f.P("PageToken: ", m.PageToken, ",")
	}
	endCall(f, t, m.Method, response)
}

type MethodSearch struct {
	Resource *annotations.ResourceDescriptor
	Method   *protogen.Method

	Parent    string
	PageSize  string
	PageToken string
}

func (m MethodSearch) Generate(f *protogen.GeneratedFile, t transport.Transport, response, err, assign string) {
	beginCall(f, t, m.Method, response, err, assign)
	if HasParent(m.Resource) {
		f.P("Parent: ", m.Parent, ",")
	}
	if m.PageSize != "" {
		f.P("PageSize: ", m.PageSize, ",")
	}
	if m.PageToken != "" {
		f.P("PageToken: ", m.PageToken, ",")
	}
	endCall(f, t, m.Method, response)
}

type MethodDelete struct {
	Resource *annotations.ResourceDescriptor
	Method   *protogen.Method

	ResourceVar string // variable name of the resource.
	Name        string
	Etag        string
}

func (m MethodDelete) Generate(f *protogen.GeneratedFile, t transport.Transport, response, err, assign string) {
	beginCall(f, t, m.Method, response, err, assign)
	if m.Name != "" {
		f.P("Name: ", m.Name, ",")
	} else {
		f.P("Name: ", m.ResourceVar, ".Name,")
	}
	switch {
	case HasEtagField(m.Method.Input.Desc) && m.Etag != "":
		f.P("Etag: ", m.Etag, ",")
	case HasRequiredEtagField(m.Method.Input.Desc):
		if m.ResourceVar != "" {
			// Delete request has an required etag field.
			f.P("Etag: ", m.ResourceVar, ".Etag,")
		} else {
			f.P("Etag: \"\",")
		}
	}
	endCall(f, t, m.Method, response)
}
