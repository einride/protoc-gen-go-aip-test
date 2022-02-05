package xrange

import (
	"google.golang.org/protobuf/reflect/protopath"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// RangeFields performs a depth-first traversal over fields in message m, calling f for each.
// It terminates when all fields have been traversed.
func RangeFields(m protoreflect.MessageDescriptor, f func(p protopath.Path, field protoreflect.FieldDescriptor)) {
	r := ranger{seen: map[protoreflect.FullName]struct{}{}}
	r.rangeMessageFields(m, protopath.Path{protopath.Root(m)}, f)
}

type ranger struct {
	seen map[protoreflect.FullName]struct{}
}

func (f *ranger) rangeMessageFields(
	message protoreflect.MessageDescriptor,
	p protopath.Path,
	fn func(p protopath.Path, field protoreflect.FieldDescriptor),
) {
	if _, ok := f.seen[message.FullName()]; ok {
		return
	}
	f.seen[message.FullName()] = struct{}{}
	fields := message.Fields()
	for i := 0; i < fields.Len(); i++ {
		field := fields.Get(i)
		// nolint: gocritic
		nextP := append(p, protopath.FieldAccess(field))
		fn(nextP, field)
		if !field.IsList() && !field.IsMap() && field.Kind() == protoreflect.MessageKind {
			f.rangeMessageFields(field.Message(), nextP, fn)
		}
	}
}
