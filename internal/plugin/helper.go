package plugin

import (
	"github.com/einride/protoc-gen-go-aiptest/internal/xrange"
	"github.com/stoewer/go-strcase"
	"go.einride.tech/aip/fieldbehavior"
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protopath"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func hasParent(resource *aipreflect.ResourceDescriptor) bool {
	if len(resource.Names) == 0 {
		return false
	}
	return len(resource.Names[0].Ancestors) > 0
}

func findMethod(service *protogen.Service, methodName protoreflect.Name) (*protogen.Method, bool) {
	for _, method := range service.Methods {
		if method.Desc.Name() == methodName {
			return method, true
		}
	}
	return nil, false
}

func hasUserSettableID(resource *aipreflect.ResourceDescriptor, method protoreflect.MethodDescriptor) bool {
	idField := strcase.LowerCamelCase(resource.Singular.UpperCamelCase()) + "_id"
	return hasField(method.Input(), protoreflect.Name(idField))
}

func hasField(message protoreflect.MessageDescriptor, field protoreflect.Name) bool {
	f := message.Fields().ByName(field)
	return f != nil
}

func hasMutableResourceReferences(message protoreflect.MessageDescriptor) bool {
	var found bool
	rangeMutableResourceReferences(message, func(p protopath.Path, field protoreflect.FieldDescriptor, r *annotations.ResourceReference) {
		found = true
	})
	return found
}

func rangeMutableResourceReferences(
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
