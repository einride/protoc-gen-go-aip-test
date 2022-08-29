package ident

import (
	"google.golang.org/protobuf/compiler/protogen"
)

// Testing idents.
//nolint: gochecknoglobals
var (
	TestingT = protogen.GoIdent{GoName: "T", GoImportPath: "testing"}
)
