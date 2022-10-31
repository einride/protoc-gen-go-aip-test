package ident

import (
	"google.golang.org/protobuf/compiler/protogen"
)

// Cmpopts idents.
//
//nolint:gochecknoglobals
var (
	CmpoptsSortSlices = protogen.GoIdent{
		GoName:       "SortSlices",
		GoImportPath: "github.com/google/go-cmp/cmp/cmpopts",
	}
)
