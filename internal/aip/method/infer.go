package method

import (
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const lro = "google.longrunning.Operation"

func inferMethod(descriptor protoreflect.MethodDescriptor) *Method {
	type method struct {
		methodType aipreflect.MethodType
		infer      func(method protoreflect.MethodDescriptor) *annotations.ResourceDescriptor
	}
	methods := []method{
		{methodType: aipreflect.MethodTypeGet, infer: getMethod},
		{methodType: aipreflect.MethodTypeBatchGet, infer: batchGetMethod},
		{methodType: aipreflect.MethodTypeCreate, infer: createMethod},
		{methodType: aipreflect.MethodTypeUpdate, infer: updateMethod},
		{methodType: aipreflect.MethodTypeList, infer: listMethod},
		{methodType: aipreflect.MethodTypeSearch, infer: searchMethod},
		{methodType: aipreflect.MethodTypeDelete, infer: deleteMethod},
	}
	for _, m := range methods {
		if resource := m.infer(descriptor); resource != nil {
			return &Method{
				Type:       m.methodType,
				Descriptor: descriptor,
				Resource:   resource,
			}
		}
	}
	return nil
}

func getMethod(method protoreflect.MethodDescriptor) *annotations.ResourceDescriptor {
	if !hasNamePrefix(method, aipreflect.MethodTypeGet) {
		return nil
	}
	resource := getResourceDescriptor(method.Output())
	return resource
}

func batchGetMethod(method protoreflect.MethodDescriptor) *annotations.ResourceDescriptor {
	if !hasNamePrefix(method, aipreflect.MethodTypeBatchGet) {
		return nil
	}
	field, resource, ok := firstFieldResource(method.Output())
	if !ok {
		return nil
	}
	if field.Cardinality() != protoreflect.Repeated {
		return nil
	}
	return resource
}

func createMethod(method protoreflect.MethodDescriptor) *annotations.ResourceDescriptor {
	if !hasNamePrefix(method, aipreflect.MethodTypeCreate) {
		return nil
	}
	output := method.Output()
	if method.Output().FullName() == lro {
		info := getLROInfo(method)
		output = findMessage(method.ParentFile(), info.GetResponseType())
	}
	resource := getResourceDescriptor(output)
	if resource == nil {
		return nil
	}
	// should have a field in input message for resource
	if findResourceField(method.Input(), resource) == nil {
		return nil
	}
	return resource
}

func updateMethod(method protoreflect.MethodDescriptor) *annotations.ResourceDescriptor {
	if !hasNamePrefix(method, aipreflect.MethodTypeUpdate) {
		return nil
	}
	output := method.Output()
	if method.Output().FullName() == lro {
		info := getLROInfo(method)
		output = findMessage(method.ParentFile(), info.GetResponseType())
	}
	resource := getResourceDescriptor(output)
	if resource == nil {
		return nil
	}
	// should have a field in input message for resource
	if findResourceField(method.Input(), resource) == nil {
		return nil
	}
	return resource
}

func listMethod(method protoreflect.MethodDescriptor) *annotations.ResourceDescriptor {
	if !hasNamePrefix(method, aipreflect.MethodTypeList) {
		return nil
	}
	output := method.Output()
	if !hasField(output, "next_page_token") {
		return nil
	}

	field, resource, ok := firstFieldResource(output)
	if !ok {
		return nil
	}
	if field.Cardinality() != protoreflect.Repeated {
		return nil
	}
	return resource
}

func searchMethod(method protoreflect.MethodDescriptor) *annotations.ResourceDescriptor {
	if !hasNamePrefix(method, aipreflect.MethodTypeList) {
		return nil
	}
	output := method.Output()
	if !hasField(output, "next_page_token") {
		return nil
	}

	field, resource, ok := firstFieldResource(output)
	if !ok {
		return nil
	}
	if field.Cardinality() != protoreflect.Repeated {
		return nil
	}
	return resource
}

func deleteMethod(method protoreflect.MethodDescriptor) *annotations.ResourceDescriptor {
	if !hasNamePrefix(method, aipreflect.MethodTypeDelete) {
		return nil
	}
	input := method.Input()
	if !hasField(input, "name") {
		return nil
	}
	nameField := input.Fields().ByName("name")
	resourceReference := getResourceReference(nameField)
	if resourceReference == nil {
		return nil
	}
	resource := findResourceDescriptor(resourceReference.GetType(), method)
	if resource == nil {
		return nil
	}
	return resource
}
