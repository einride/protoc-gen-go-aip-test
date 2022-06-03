package util

import (
	"strings"

	"github.com/einride/protoc-gen-go-aip-test/internal/aip/method"
	"github.com/einride/protoc-gen-go-aip-test/internal/xrange"
	"github.com/stoewer/go-strcase"
	"go.einride.tech/aip/fieldbehavior"
	"go.einride.tech/aip/resourcename"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
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

func WildcardResourceName(r *annotations.ResourceDescriptor) string {
	patternSgments := resourceNameSegments(r.GetPattern()[0])
	nameSegments := make([]string, 0, len(patternSgments))
	for _, segment := range patternSgments {
		if segment.IsVariable() {
			nameSegments = append(nameSegments, resourcename.Wildcard)
		} else {
			nameSegments = append(nameSegments, string(segment))
		}
	}
	return strings.Join(nameSegments, "/")
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

func HasAnyStandardMethodFor(s protoreflect.ServiceDescriptor, r *annotations.ResourceDescriptor) bool {
	for _, resource := range method.NewMethods(s).Resources() {
		if resource.Type == r.Type {
			return true
		}
	}
	return false
}

func ReturnsLRO(method protoreflect.MethodDescriptor) bool {
	return method.Output().FullName() == "google.longrunning.Operation"
}

func IsAlternativeBatch(method protoreflect.MethodDescriptor) bool {
	switch {
	case strings.HasPrefix(string(method.Name()), "BatchGet"):
		return IsAlternativeBatchGet(method)
	case strings.HasPrefix(string(method.Name()), "BatchDelete"):
		return IsAlternativeBatchDelete(method)
	default:
		return false
	}
}

func IsAlternativeBatchGet(method protoreflect.MethodDescriptor) bool {
	if !strings.HasPrefix(string(method.Name()), "BatchGet") {
		return false
	}
	inputFields := method.Input().Fields()
	return inputFields.ByName("requests") != nil
}

func IsAlternativeBatchDelete(method protoreflect.MethodDescriptor) bool {
	if !strings.HasPrefix(string(method.Name()), "BatchDelete") {
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

func FindResourceField(
	message protoreflect.MessageDescriptor,
	resource *annotations.ResourceDescriptor,
) protoreflect.FieldDescriptor {
	for i := 0; i < message.Fields().Len(); i++ {
		field := message.Fields().Get(i)
		if field.Kind() == protoreflect.MessageKind {
			r := getResourceDescriptor(field.Message())
			if r != nil && r.Type == resource.Type {
				return field
			}
		}
	}
	return nil
}

func getResourceDescriptor(message protoreflect.MessageDescriptor) *annotations.ResourceDescriptor {
	return proto.GetExtension(
		message.Options(),
		annotations.E_Resource,
	).(*annotations.ResourceDescriptor)
}
