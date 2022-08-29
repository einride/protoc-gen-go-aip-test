package ident

import (
	"google.golang.org/protobuf/compiler/protogen"
)

// Proto idents.
//nolint: gochecknoglobals
var (
	ProtoClone = protogen.GoIdent{
		GoName:       "Clone",
		GoImportPath: "google.golang.org/protobuf/proto",
	}
)
