package util

import (
	"github.com/stoewer/go-strcase"
)

// APIMode represents the protobuf API mode for generated code.
type APIMode int

const (
	// APIModeOpen generates code for the Open Struct API (default).
	APIModeOpen APIMode = iota
	// APIModeOpaque generates code for the Opaque API.
	APIModeOpaque
)

// FieldGet generates code for accessing a field value.
// For Open Struct API: varName.FieldName
// For Opaque API: varName.GetFieldName().
func FieldGet(varName, fieldName string, apiMode APIMode) string {
	if apiMode == APIModeOpaque {
		return varName + ".Get" + strcase.UpperCamelCase(fieldName) + "()"
	}
	return varName + "." + strcase.UpperCamelCase(fieldName)
}

// FieldSet generates code for setting a field value.
// For Open Struct API: varName.FieldName = value
// For Opaque API: varName.SetFieldName(value).
func FieldSet(varName, fieldName, value string, apiMode APIMode) string {
	if apiMode == APIModeOpaque {
		return varName + ".Set" + strcase.UpperCamelCase(fieldName) + "(" + value + ")"
	}
	return varName + "." + strcase.UpperCamelCase(fieldName) + " = " + value
}
