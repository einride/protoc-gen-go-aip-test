package ident

import (
	"google.golang.org/protobuf/compiler/protogen"
)

var (
	TimeSecond = protogen.GoIdent{GoName: "Second", GoImportPath: "time"}
	TimeSince  = protogen.GoIdent{GoName: "Since", GoImportPath: "time"}
)
