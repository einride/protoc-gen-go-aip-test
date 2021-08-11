package plugin

import (
	"github.com/stoewer/go-strcase"
	"go.einride.tech/aip/reflect/aipreflect"
	"go.einride.tech/aip/resourcename"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func resourceType(r *annotations.ResourceDescriptor) string {
	return aipreflect.ResourceType(r.GetType()).Type()
}

func hasUserSettableIDField(r *annotations.ResourceDescriptor, m protoreflect.MessageDescriptor) bool {
	idField := strcase.LowerCamelCase(r.GetSingular()) + "_id"
	return m.Fields().ByName(protoreflect.Name(idField)) != nil
}

func hasParent(r *annotations.ResourceDescriptor) bool {
	return len(resourceNameSegments(r.GetPattern()[0])) > 3
}

func resourceNameSegments(pattern string) []resourcename.Segment {
	var s resourcename.Scanner
	s.Init(pattern)
	segments := make([]resourcename.Segment, 0)
	for s.Scan() {
		segments = append(segments, s.Segment())
	}
	return segments
}

func allStandardMethods() []aipreflect.MethodType {
	return []aipreflect.MethodType{
		aipreflect.MethodTypeGet,
		aipreflect.MethodTypeList,
		aipreflect.MethodTypeCreate,
		aipreflect.MethodTypeUpdate,
		aipreflect.MethodTypeDelete,
		aipreflect.MethodTypeUndelete,
		aipreflect.MethodTypeBatchGet,
		aipreflect.MethodTypeBatchCreate,
		aipreflect.MethodTypeBatchUpdate,
		aipreflect.MethodTypeBatchDelete,
		aipreflect.MethodTypeSearch,
	}
}

func hasAnyStandardMethodFor(s protoreflect.ServiceDescriptor, r *annotations.ResourceDescriptor) bool {
	methods := resourceStandardMethods(r)
	for _, method := range methods {
		if s.Methods().ByName(method) != nil {
			return true
		}
	}
	return false
}

func resourceStandardMethods(r *annotations.ResourceDescriptor) []protoreflect.Name {
	methodTypes := allStandardMethods()
	standardMethods := make([]protoreflect.Name, 0, len(methodTypes))
	for _, methodType := range methodTypes {
		standardMethods = append(standardMethods, inferMethodName(r, methodType))
	}
	return standardMethods
}

func inferMethodName(r *annotations.ResourceDescriptor, methodType aipreflect.MethodType) protoreflect.Name {
	grammaticalName := aipreflect.GrammaticalName(r.GetSingular())
	if methodType.IsPlural() {
		grammaticalName = aipreflect.GrammaticalName(r.GetPlural())
	}
	return methodType.NamePrefix() + protoreflect.Name(grammaticalName.UpperCamelCase())
}
