package ident

import (
	"google.golang.org/protobuf/compiler/protogen"
)

// Strings idents.
//
//nolint:gochecknoglobals
var (
	StringsHasSuffix = protogen.GoIdent{GoName: "HasSuffix", GoImportPath: "strings"}
)
