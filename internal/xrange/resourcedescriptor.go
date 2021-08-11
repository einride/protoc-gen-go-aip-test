package xrange

import (
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// RangeResourceDescriptors traverses over resource descriptors in file, calling f for each.
// If the resource descriptor does not belong to a message, the corresponding parameter will be nil.
func RangeResourceDescriptors(
	file protoreflect.FileDescriptor,
	f func(protoreflect.MessageDescriptor, *annotations.ResourceDescriptor),
) {
	forwardedDescriptors := proto.GetExtension(
		file.Options(),
		annotations.E_ResourceDefinition,
	).([]*annotations.ResourceDescriptor)
	for _, descriptor := range forwardedDescriptors {
		f(nil, descriptor)
	}

	for i := 0; i < file.Messages().Len(); i++ {
		message := file.Messages().Get(i)
		descriptor := proto.GetExtension(
			message.Options(),
			annotations.E_Resource,
		).(*annotations.ResourceDescriptor)
		if descriptor == nil {
			continue
		}
		f(message, descriptor)
	}
}
