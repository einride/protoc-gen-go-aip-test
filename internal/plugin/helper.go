package plugin

import (
	"github.com/einride/protoc-gen-go-aiptest/internal/xrange"
	"go.einride.tech/aip/fieldbehavior"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/reflect/protopath"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func hasMutableResourceReferences(message protoreflect.MessageDescriptor) bool {
	var found bool
	rangeMutableResourceReferences(
		message,
		func(p protopath.Path, field protoreflect.FieldDescriptor, r *annotations.ResourceReference) {
			found = true
		},
	)
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

func hasRequiredFields(message protoreflect.MessageDescriptor) bool {
	var found bool
	rangeRequiredFields(message, func(p protopath.Path, field protoreflect.FieldDescriptor) {
		found = true
	})
	return found
}

func rangeRequiredFields(
	message protoreflect.MessageDescriptor,
	f func(protopath.Path, protoreflect.FieldDescriptor),
) {
	xrange.RangeFields(
		message,
		func(p protopath.Path, field protoreflect.FieldDescriptor) {
			if fieldbehavior.Has(field, annotations.FieldBehavior_REQUIRED) {
				f(p, field)
			}
		},
	)
}
