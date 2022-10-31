package ident

import (
	"google.golang.org/protobuf/compiler/protogen"
)

// Assert idents.
//
//nolint:gochecknoglobals
var (
	AssertEqual = protogen.GoIdent{
		GoName:       "Equal",
		GoImportPath: "gotest.tools/v3/assert",
	}

	AssertDeepEqual = protogen.GoIdent{
		GoName:       "DeepEqual",
		GoImportPath: "gotest.tools/v3/assert",
	}

	AssertNilError = protogen.GoIdent{
		GoName:       "NilError",
		GoImportPath: "gotest.tools/v3/assert",
	}

	AssertCheck = protogen.GoIdent{
		GoName:       "Check",
		GoImportPath: "gotest.tools/v3/assert",
	}
)
