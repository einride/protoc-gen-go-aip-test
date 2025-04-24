package ident

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/compiler/protogen"
)

func Codes(code codes.Code) protogen.GoIdent {
	return protogen.GoIdent{
		GoName:       "Code" + code.String(),
		GoImportPath: "connectrpc.com/connect",
	}
}
