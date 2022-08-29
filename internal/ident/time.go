package ident

import (
	"google.golang.org/protobuf/compiler/protogen"
)

// Time idents.
//nolint: gochecknoglobals
var (
	TimeSecond = protogen.GoIdent{GoName: "Second", GoImportPath: "time"}
	TimeSince  = protogen.GoIdent{GoName: "Since", GoImportPath: "time"}
)
