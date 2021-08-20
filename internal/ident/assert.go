package ident

import (
	"google.golang.org/protobuf/compiler/protogen"
)

var AssertEqual = protogen.GoIdent{
	GoName:       "Equal",
	GoImportPath: "gotest.tools/v3/assert",
}

var AssertDeepEqual = protogen.GoIdent{
	GoName:       "DeepEqual",
	GoImportPath: "gotest.tools/v3/assert",
}

var AssertNilError = protogen.GoIdent{
	GoName:       "NilError",
	GoImportPath: "gotest.tools/v3/assert",
}

var AssertCheck = protogen.GoIdent{
	GoName:       "Check",
	GoImportPath: "gotest.tools/v3/assert",
}
