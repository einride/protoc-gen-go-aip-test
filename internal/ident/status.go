package ident

import (
	"google.golang.org/protobuf/compiler/protogen"
)

// Status idents.
//
//nolint:gochecknoglobals
var (
	StatusCode = protogen.GoIdent{
		GoName:       "CodeOf",
		GoImportPath: "connectrpc.com/connect",
	}
)
