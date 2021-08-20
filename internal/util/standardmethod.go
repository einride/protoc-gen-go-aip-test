package util

import (
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
)

func StandardMethod(
	service *protogen.Service,
	r *annotations.ResourceDescriptor,
	methodType aipreflect.MethodType,
) (*protogen.Method, bool) {
	methodName := InferMethodName(r, methodType)
	for _, method := range service.Methods {
		if method.Desc.Name() == methodName {
			return method, true
		}
	}
	return nil, false
}
