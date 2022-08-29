package ident

import (
	"google.golang.org/protobuf/compiler/protogen"
)

// Protocmp idents.
//nolint: gochecknoglobals
var (
	ProtocmpTransform = protogen.GoIdent{
		GoName:       "Transform",
		GoImportPath: "google.golang.org/protobuf/testing/protocmp",
	}
)
