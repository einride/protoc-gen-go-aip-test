package plugin

import (
	"github.com/stoewer/go-strcase"
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/protobuf/compiler/protogen"
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
