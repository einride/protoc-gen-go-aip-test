package ident

import (
	"google.golang.org/protobuf/compiler/protogen"
)

var StatusCode = protogen.GoIdent{
	GoName:       "Code",
	GoImportPath: "google.golang.org/grpc/status",
}
