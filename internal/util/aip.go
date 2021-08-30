package util

import (
	"strings"

	"github.com/einride/protoc-gen-go-aip-test/internal/xrange"
	"github.com/stoewer/go-strcase"
	"go.einride.tech/aip/fieldbehavior"
	"go.einride.tech/aip/reflect/aipreflect"
	"go.einride.tech/aip/resourcename"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/reflect/protopath"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func HasUserSettableIDField(r *annotations.ResourceDescriptor, m protoreflect.MessageDescriptor) bool {
	idField := strcase.LowerCamelCase(r.GetSingular()) + "_id"
	return m.Fields().ByName(protoreflect.Name(idField)) != nil
}

func HasParent(r *annotations.ResourceDescriptor) bool {
	return len(resourceNameSegments(r.GetPattern()[0])) > 3
}

func resourceNameSegments(pattern string) []resourcename.Segment {
	var s resourcename.Scanner
	s.Init(pattern)
	segments := make([]resourcename.Segment, 0)
	for s.Scan() {
		segments = append(segments, s.Segment())
	}
	return segments
}

func allStandardMethods() []aipreflect.MethodType {
	return []aipreflect.MethodType{
		aipreflect.MethodTypeGet,
		aipreflect.MethodTypeList,
		aipreflect.MethodTypeCreate,
		aipreflect.MethodTypeUpdate,
		aipreflect.MethodTypeDelete,
		aipreflect.MethodTypeUndelete,
		aipreflect.MethodTypeBatchGet,
		aipreflect.MethodTypeBatchCreate,
		aipreflect.MethodTypeBatchUpdate,
		aipreflect.MethodTypeBatchDelete,
		aipreflect.MethodTypeSearch,
	}
}

func HasAnyStandardMethodFor(s protoreflect.ServiceDescriptor, r *annotations.ResourceDescriptor) bool {
	methods := ResourceStandardMethods(r)
	for _, method := range methods {
		if s.Methods().ByName(method) != nil {
			return true
		}
	}
	return false
}

func ResourceStandardMethods(r *annotations.ResourceDescriptor) []protoreflect.Name {
	methodTypes := allStandardMethods()
	standardMethods := make([]protoreflect.Name, 0, len(methodTypes))
	for _, methodType := range methodTypes {
		standardMethods = append(standardMethods, InferMethodName(r, methodType))
	}
	return standardMethods
}

func InferMethodName(r *annotations.ResourceDescriptor, methodType aipreflect.MethodType) protoreflect.Name {
	grammaticalName := aipreflect.GrammaticalName(r.GetSingular())
	if methodType.IsPlural() {
		grammaticalName = aipreflect.GrammaticalName(r.GetPlural())
	}
	return methodType.NamePrefix() + protoreflect.Name(grammaticalName.UpperCamelCase())
}

func ReturnsLRO(method protoreflect.MethodDescriptor) bool {
	return method.Output().FullName() == "google.longrunning.Operation"
}

func IsAlternativeBatchGet(method protoreflect.MethodDescriptor) bool {
	if !strings.HasPrefix(string(method.Name()), "BatchGet") {
		return false
	}
	inputFields := method.Input().Fields()
	return inputFields.ByName("requests") != nil
}

func HasUpdateMask(method protoreflect.MethodDescriptor) bool {
	if !strings.HasPrefix(string(method.Name()), "Update") {
		return false
	}
	return method.Input().Fields().ByName("update_mask") != nil
}

func HasRequiredFields(message protoreflect.MessageDescriptor) bool {
	var found bool
	RangeRequiredFields(message, func(p protopath.Path, field protoreflect.FieldDescriptor) {
		found = true
	})
	return found
}

func RangeRequiredFields(
	message protoreflect.MessageDescriptor,
	f func(protopath.Path, protoreflect.FieldDescriptor),
) {
	xrange.RangeFields(
		message,
		func(p protopath.Path, field protoreflect.FieldDescriptor) {
			if fieldbehavior.Has(field, annotations.FieldBehavior_REQUIRED) {
				f(p, field)
			}
		},
	)
}

func HasMutableResourceReferences(message protoreflect.MessageDescriptor) bool {
	var found bool
	RangeMutableResourceReferences(
		message,
		func(p protopath.Path, field protoreflect.FieldDescriptor, r *annotations.ResourceReference) {
			found = true
		},
	)
	return found
}

func RangeMutableResourceReferences(
	message protoreflect.MessageDescriptor,
	f func(protopath.Path, protoreflect.FieldDescriptor, *annotations.ResourceReference),
) {
	xrange.RangeResourceReferences(
		message,
		func(p protopath.Path, field protoreflect.FieldDescriptor, r *annotations.ResourceReference) {
			if fieldbehavior.Has(field, annotations.FieldBehavior_OUTPUT_ONLY) {
				return
			}
			f(p, field, r)
		},
	)
}

func PathChainGet(p protopath.Path) string {
	gg := make([]string, 0, len(p))
	for _, step := range p {
		g := "Get" + strcase.UpperCamelCase(string(step.FieldDescriptor().Name())) + "()"
		gg = append(gg, g)
	}
	return strings.Join(gg, ".")
}
