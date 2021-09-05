package method

import (
	"strings"

	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/genproto/googleapis/api/annotations"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func getResourceDescriptor(message protoreflect.MessageDescriptor) *annotations.ResourceDescriptor {
	return proto.GetExtension(
		message.Options(),
		annotations.E_Resource,
	).(*annotations.ResourceDescriptor)
}

func getResourceReference(field protoreflect.FieldDescriptor) *annotations.ResourceReference {
	return proto.GetExtension(
		field.Options(),
		annotations.E_ResourceReference,
	).(*annotations.ResourceReference)
}

func hasNamePrefix(method protoreflect.Descriptor, methodType aipreflect.MethodType) bool {
	return strings.HasPrefix(string(method.Name()), methodType.String())
}

func findMessage(file protoreflect.FileDescriptor, name string) protoreflect.MessageDescriptor {
	var message protoreflect.MessageDescriptor
	rangeReachableMessages(file, func(m protoreflect.MessageDescriptor) bool {
		if string(m.FullName()) == name || string(m.Name()) == name {
			message = m
			return false
		}
		return true
	})
	return message
}

func getLROInfo(method protoreflect.MethodDescriptor) *longrunningpb.OperationInfo {
	return proto.GetExtension(
		method.Options(),
		longrunningpb.E_OperationInfo,
	).(*longrunningpb.OperationInfo)
}

func rangeReachableMessages(file protoreflect.FileDescriptor, fn func(message protoreflect.MessageDescriptor) bool) {
	if !rangeFileMessages(file, fn) {
		return
	}
	for i := 0; i < file.Imports().Len(); i++ {
		if !rangeFileMessages(file.Imports().Get(i).FileDescriptor, fn) {
			return
		}
	}
}

func rangeFileMessages(file protoreflect.FileDescriptor, fn func(message protoreflect.MessageDescriptor) bool) bool {
	for i := 0; i < file.Messages().Len(); i++ {
		message := file.Messages().Get(i)
		if !fn(message) {
			return false
		}
		if !rangeMessageMessages(message, fn) {
			return false
		}
	}
	return true
}

func rangeMessageMessages(
	message protoreflect.MessageDescriptor,
	fn func(message protoreflect.MessageDescriptor) bool,
) bool {
	for i := 0; i < message.Messages().Len(); i++ {
		nested := message.Messages().Get(i)
		if !fn(nested) {
			return false
		}
		if !rangeMessageMessages(nested, fn) {
			return false
		}
	}
	return true
}

func findResourceField(
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

func firstFieldResource(
	message protoreflect.MessageDescriptor,
) (protoreflect.FieldDescriptor, *annotations.ResourceDescriptor, bool) {
	for i := 0; i < message.Fields().Len(); i++ {
		field := message.Fields().Get(i)
		if field.Kind() == protoreflect.MessageKind {
			r := getResourceDescriptor(field.Message())
			if r != nil {
				return field, r, true
			}
		}
	}
	return nil, nil, false
}

func hasField(message protoreflect.MessageDescriptor, name string) bool {
	return message.Fields().ByName(protoreflect.Name(name)) != nil
}

func findResourceDescriptor(typ string, scope protoreflect.Descriptor) *annotations.ResourceDescriptor {
	var resourceDescriptor *annotations.ResourceDescriptor
	rangeReachableFiles(scope, func(file protoreflect.FileDescriptor) bool {
		if resourceDescriptor != nil {
			return false
		}
		aipreflect.RangeResourceDescriptorsInFile(file, func(resource *annotations.ResourceDescriptor) bool {
			if resource.Type == typ {
				resourceDescriptor = resource
				return false
			}
			return true
		})
		return true
	})
	return resourceDescriptor
}

func rangeReachableFiles(desc protoreflect.Descriptor, fn func(file protoreflect.FileDescriptor) bool) {
	if !fn(desc.ParentFile()) {
		return
	}
	for i := 0; i < desc.ParentFile().Imports().Len(); i++ {
		if !fn(desc.ParentFile().Imports().Get(i).FileDescriptor) {
			return
		}
	}
}
