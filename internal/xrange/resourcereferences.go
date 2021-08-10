package xrange

import (
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protopath"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// RangeResourceReferences performs a depth-first traversal over fields in message m, calling f for each
// field with a resource reference annotation.
// It terminates when all fields have been traversed.
func RangeResourceReferences(
	m protoreflect.MessageDescriptor,
	f func(protopath.Path, protoreflect.FieldDescriptor, *annotations.ResourceReference),
) {
	RangeFields(m, func(p protopath.Path, field protoreflect.FieldDescriptor) {
		resourceReference, ok := getResourceReference(field)
		if ok {
			f(p, field, resourceReference)
		}
	})
}

func getResourceReference(field protoreflect.FieldDescriptor) (*annotations.ResourceReference, bool) {
	r := proto.GetExtension(field.Options(), annotations.E_ResourceReference).(*annotations.ResourceReference)
	if r == nil {
		return nil, false
	}
	return r, true
}
