package util

import (
	"github.com/einride/protoc-gen-go-aip-test/internal/aip/method"
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
)

func StandardMethod(
	service *protogen.Service,
	r *annotations.ResourceDescriptor,
	methodType aipreflect.MethodType,
) (*protogen.Method, bool) {
	methods := method.NewMethods(service.Desc)
	m := methods.Get(r, methodType)
	if m == nil {
		return nil, false
	}

	methodName := m.Descriptor.Name()
	for _, method := range service.Methods {
		if method.Desc.Name() == methodName {
			return method, true
		}
	}
	return nil, false
}
