package ident

import (
	"google.golang.org/protobuf/compiler/protogen"
)

// Status idents.
//
//nolint:gochecknoglobals
var (
	StatusCode = protogen.GoIdent{GoName: "Code", GoImportPath: "google.golang.org/grpc/status"}
)
