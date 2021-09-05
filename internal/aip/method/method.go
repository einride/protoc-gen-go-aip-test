package method

import (
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// Methods represents a set of Method, usually found on a
// protoreflect.ServiceDescriptor.
type Methods []*Method

// Method represents one of the AIP standard methods for a
// specific resource.
type Method struct {
	Type       aipreflect.MethodType
	Descriptor protoreflect.MethodDescriptor
	Resource   *annotations.ResourceDescriptor
}

func NewMethods(service protoreflect.ServiceDescriptor) Methods {
	methods := make([]*Method, 0, service.Methods().Len())
	for i := 0; i < service.Methods().Len(); i++ {
		if method := inferMethod(service.Methods().Get(i)); method != nil {
			methods = append(methods, method)
		}
	}
	return methods
}

// Resources returns all unique resources Methods refers to.
func (m Methods) Resources() []*annotations.ResourceDescriptor {
	seen := make(map[string]struct{})
	resources := make([]*annotations.ResourceDescriptor, 0, len(m))
	for _, method := range m {
		if _, ok := seen[method.Resource.Type]; ok {
			continue
		}
		seen[method.Resource.Type] = struct{}{}
		resources = append(resources, method.Resource)
	}
	return resources
}

// Get returns the Method with methodType for resource.
// If no such method exists, nil is returned.
func (m Methods) Get(
	resource *annotations.ResourceDescriptor,
	methodType aipreflect.MethodType,
) *Method {
	for _, method := range m {
		if method.Type == methodType && method.Resource.Type == resource.Type {
			return method
		}
	}
	return nil
}
