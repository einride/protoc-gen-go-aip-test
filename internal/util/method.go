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

func (m MethodCreate) Generate(f *protogen.GeneratedFile, request, response, err, assign string, apiMode APIMode) {
	userSetID := m.UserSettableID
	if userSetID == "" && HasUserSettableIDField(m.Resource, m.Method.Input.Desc) {
		userSetID = "userSetID"
		f.P(userSetID + " := \"\"")
		f.P("if fx.IDGenerator != nil {")
		f.P(userSetID + " = fx.IDGenerator()")
		f.P("}")
	}

	upper := strcase.UpperCamelCase(string(FindResourceField(
		m.Method.Input.Desc,
		m.Resource,
	).Name()))

	if apiMode == APIModeOpaque {
		// Opaque API: create request, set fields, then call
		f.P(request, " := &", m.Method.Input.GoIdent, "{}")
		if HasParent(m.Resource) {
			f.P(request, ".SetParent(", m.Parent, ")")
		}
		switch {
		case m.Message != "":
			f.P(request, ".Set", upper, "(", m.Message, ")")
		case !HasParent(m.Resource):
			f.P(request, ".Set", upper, "(fx.Create())")
		default:
			f.P(request, ".Set", upper, "(fx.Create(", m.Parent, "))")
		}
		if userSetID != "" && HasUserSettableIDField(m.Resource, m.Method.Input.Desc) {
			f.P(request, ".Set", upper, "Id(", userSetID, ")")
		}
		f.P(response, ", ", err, " ", assign, " fx.Service().", m.Method.GoName, "(fx.Context(), ", request, ")")
	} else {
		// Open Struct API: inline struct literal
		f.P(
			response,
			", ",
			err,
			" ",
			assign,
			" fx.Service().",
			m.Method.GoName,
			"(fx.Context(), &",
			m.Method.Input.GoIdent,
			"{",
		)
		if HasParent(m.Resource) {
			f.P("Parent: ", m.Parent, ",")
		}
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
		f.P("})")
	}
}

type MethodGet struct {
	Resource *annotations.ResourceDescriptor
	Method   *protogen.Method

	Name string
}

func (m MethodGet) Generate(f *protogen.GeneratedFile, request, response, err, assign string, apiMode APIMode) {
	if apiMode == APIModeOpaque {
		f.P(request, " := &", m.Method.Input.GoIdent, "{}")
		f.P(request, ".SetName(", m.Name, ")")
		f.P(response, ", ", err, " ", assign, " fx.Service().", m.Method.GoName, "(fx.Context(), ", request, ")")
	} else {
		f.P(
			response,
			", ",
			err,
			" ",
			assign,
			" fx.Service().",
			m.Method.GoName,
			"(fx.Context(), &",
			m.Method.Input.GoIdent,
			"{",
		)
		f.P("Name: ", m.Name, ",")
		f.P("})")
	}
}

type MethodBatchGet struct {
	Resource *annotations.ResourceDescriptor
	Method   *protogen.Method

	Parent string
	Names  []string
}

func (m MethodBatchGet) Generate(f *protogen.GeneratedFile, request, response, err, assign string, apiMode APIMode) {
	if apiMode == APIModeOpaque {
		f.P(request, " := &", m.Method.Input.GoIdent, "{}")
		if HasParent(m.Resource) {
			f.P(request, ".SetParent(", m.Parent, ")")
		}
		f.P(request, ".SetNames([]string{")
		for _, name := range m.Names {
			f.P(name, ",")
		}
		f.P("})")
		f.P(response, ", ", err, " ", assign, " fx.Service().", m.Method.GoName, "(fx.Context(), ", request, ")")
	} else {
		f.P(
			response,
			", ",
			err,
			" ",
			assign,
			" fx.Service().",
			m.Method.GoName,
			"(fx.Context(), &",
			m.Method.Input.GoIdent,
			"{",
		)
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
}

type MethodUpdate struct {
	Resource *annotations.ResourceDescriptor
	Method   *protogen.Method

	// set either Parent + Name, or Msg
	Name       string
	Parent     string
	Msg        string
	UpdateMask []string
	Etag       string
	EtagTest   bool
}

func (m MethodUpdate) Generate(f *protogen.GeneratedFile, request, response, err, assign string, apiMode APIMode) {
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
		if apiMode == APIModeOpaque {
			f.P("msg.SetName(", m.Name, ")")
		} else {
			f.P("msg.Name = ", m.Name)
		}
	}
	if m.EtagTest && !HasEtagField(m.Method.Input.Desc) && HasEtagField(m.Method.Output.Desc) {
		// Request object does not have an etag field, but the resource has.
		if apiMode == APIModeOpaque {
			if m.Etag != "" {
				f.P("msg.SetEtag(", m.Etag, ")")
			} else {
				f.P("msg.SetEtag(created.GetEtag()) // assign etag from the created resource")
			}
		} else {
			if m.Etag != "" {
				f.P("msg.Etag = ", m.Etag)
			} else {
				f.P(`msg.Etag = created.Etag // assign etag from the created resource`)
			}
		}
	}

	if apiMode == APIModeOpaque {
		f.P(request, " := &", m.Method.Input.GoIdent, "{}")
		if m.Msg != "" {
			f.P(request, ".Set", upper, "(", m.Msg, ")")
		} else {
			f.P(request, ".Set", upper, "(msg)")
		}
		if HasUpdateMask(m.Method.Desc) && len(m.UpdateMask) > 0 {
			fieldmaskpbFieldMask := f.QualifiedGoIdent(protogen.GoIdent{
				GoName:       "FieldMask",
				GoImportPath: "google.golang.org/protobuf/types/known/fieldmaskpb",
			})
			f.P(request, ".SetUpdateMask(&", fieldmaskpbFieldMask, "{")
			f.P("Paths: []string{")
			for _, path := range m.UpdateMask {
				f.P(path, ",")
			}
			f.P("},")
			f.P("})")
		}
		switch {
		case HasEtagField(m.Method.Input.Desc) && m.Etag != "":
			f.P(request, ".SetEtag(", m.Etag, ")")
		case HasRequiredEtagField(m.Method.Input.Desc):
			if m.Msg != "" {
				f.P(request, ".SetEtag(", m.Msg, ".GetEtag())")
			} else {
				f.P(request, ".SetEtag(msg.GetEtag())")
			}
		}
		f.P(response, ", ", err, " ", assign, " fx.Service().", m.Method.GoName, "(fx.Context(), ", request, ")")
	} else {
		f.P(
			response,
			", ",
			err,
			" ",
			assign,
			" fx.Service().",
			m.Method.GoName,
			"(fx.Context(), &",
			m.Method.Input.GoIdent,
			"{",
		)
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
		f.P("})")
	}
}

type MethodList struct {
	Resource *annotations.ResourceDescriptor
	Method   *protogen.Method

	Parent    string
	PageSize  string
	PageToken string
}

func (m MethodList) Generate(f *protogen.GeneratedFile, request, response, err, assign string, apiMode APIMode) {
	if apiMode == APIModeOpaque {
		f.P(request, " := &", m.Method.Input.GoIdent, "{}")
		if HasParent(m.Resource) {
			f.P(request, ".SetParent(", m.Parent, ")")
		}
		if m.PageSize != "" {
			f.P(request, ".SetPageSize(", m.PageSize, ")")
		}
		if m.PageToken != "" {
			f.P(request, ".SetPageToken(", m.PageToken, ")")
		}
		f.P(response, ", ", err, " ", assign, " fx.Service().", m.Method.GoName, "(fx.Context(), ", request, ")")
	} else {
		f.P(
			response,
			", ",
			err,
			" ",
			assign,
			" fx.Service().",
			m.Method.GoName,
			"(fx.Context(), &",
			m.Method.Input.GoIdent,
			"{",
		)
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
}

type MethodSearch struct {
	Resource *annotations.ResourceDescriptor
	Method   *protogen.Method

	Parent    string
	PageSize  string
	PageToken string
}

func (m MethodSearch) Generate(f *protogen.GeneratedFile, request, response, err, assign string, apiMode APIMode) {
	if apiMode == APIModeOpaque {
		f.P(request, " := &", m.Method.Input.GoIdent, "{}")
		if HasParent(m.Resource) {
			f.P(request, ".SetParent(", m.Parent, ")")
		}
		if m.PageSize != "" {
			f.P(request, ".SetPageSize(", m.PageSize, ")")
		}
		if m.PageToken != "" {
			f.P(request, ".SetPageToken(", m.PageToken, ")")
		}
		f.P(response, ", ", err, " ", assign, " fx.Service().", m.Method.GoName, "(fx.Context(), ", request, ")")
	} else {
		f.P(
			response,
			", ",
			err,
			" ",
			assign,
			" fx.Service().",
			m.Method.GoName,
			"(fx.Context(), &",
			m.Method.Input.GoIdent,
			"{",
		)
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
}

type MethodDelete struct {
	Resource *annotations.ResourceDescriptor
	Method   *protogen.Method

	ResourceVar string // variable name of the resource.
	Name        string
	Etag        string
}

func (m MethodDelete) Generate(f *protogen.GeneratedFile, request, response, err, assign string, apiMode APIMode) {
	if apiMode == APIModeOpaque {
		f.P(request, " := &", m.Method.Input.GoIdent, "{}")
		if m.Name != "" {
			f.P(request, ".SetName(", m.Name, ")")
		} else {
			f.P(request, ".SetName(", m.ResourceVar, ".GetName())")
		}
		switch {
		case HasEtagField(m.Method.Input.Desc) && m.Etag != "":
			f.P(request, ".SetEtag(", m.Etag, ")")
		case HasRequiredEtagField(m.Method.Input.Desc):
			if m.ResourceVar != "" {
				f.P(request, ".SetEtag(", m.ResourceVar, ".GetEtag())")
			} else {
				f.P(request, `.SetEtag("")`)
			}
		}
		f.P(response, ", ", err, " ", assign, " fx.Service().", m.Method.GoName, "(fx.Context(), ", request, ")")
	} else {
		f.P(
			response,
			", ",
			err,
			" ",
			assign,
			" fx.Service().",
			m.Method.GoName,
			"(fx.Context(), &",
			m.Method.Input.GoIdent,
			"{",
		)
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
		f.P("})")
	}
}
