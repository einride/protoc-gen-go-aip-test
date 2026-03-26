package transport

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/compiler/protogen"
)

// Transport represents the target RPC framework for generated test code.
type Transport string

const (
	// GRPC generates tests targeting protoc-gen-go-grpc's ServiceServer interface.
	GRPC Transport = "grpc"
	// Connect generates tests targeting protoc-gen-connect-go's ServiceHandler interface
	// with connect.Request[T]/connect.Response[T] wrappers.
	Connect Transport = "connect"
	// ConnectSimple generates tests targeting protoc-gen-connect-go's ServiceHandler interface
	// in simple mode (bare signatures matching gRPC).
	ConnectSimple Transport = "connect-simple"
)

const connectImportPath = "connectrpc.com/connect"

// Parse parses a transport option string. An empty string defaults to GRPC.
func Parse(s string) (Transport, error) {
	switch Transport(s) {
	case "", GRPC:
		return GRPC, nil
	case Connect:
		return Connect, nil
	case ConnectSimple:
		return ConnectSimple, nil
	default:
		return "", fmt.Errorf("unknown transport %q (expected grpc, connect, or connect-simple)", s)
	}
}

// IsConnect returns true for both connect and connect-simple transports.
func (t Transport) IsConnect() bool {
	return t == Connect || t == ConnectSimple
}

// UsesRequestWrapper returns true only for standard connect mode,
// which wraps requests in connect.Request[T] and responses in connect.Response[T].
func (t Transport) UsesRequestWrapper() bool {
	return t == Connect
}

// ServiceIdent returns the GoIdent for the service interface type.
//
// For gRPC: {ServiceName}Server from the proto package.
// For connect: {ServiceName}Handler from the {proto_pkg}connect sub-package.
func (t Transport) ServiceIdent(
	service *protogen.Service,
	goPackageName protogen.GoPackageName,
) protogen.GoIdent {
	protoImportPath := service.Methods[0].Input.GoIdent.GoImportPath
	if t.IsConnect() {
		return protogen.GoIdent{
			GoName:       service.GoName + "Handler",
			GoImportPath: protogen.GoImportPath(string(protoImportPath) + "/" + string(goPackageName) + "connect"),
		}
	}
	return protogen.GoIdent{
		GoName:       service.GoName + "Server",
		GoImportPath: protoImportPath,
	}
}

// CodesIdent returns the GoIdent for an error code constant.
//
// For gRPC: codes.NotFound from google.golang.org/grpc/codes.
// For connect: connect.CodeNotFound from connectrpc.com/connect.
func (t Transport) CodesIdent(code codes.Code) protogen.GoIdent {
	if t.IsConnect() {
		return protogen.GoIdent{
			GoName:       "Code" + code.String(),
			GoImportPath: connectImportPath,
		}
	}
	return protogen.GoIdent{
		GoName:       code.String(),
		GoImportPath: "google.golang.org/grpc/codes",
	}
}

// CodeOfIdent returns the GoIdent for the function that extracts an error code from an error.
//
// For gRPC: status.Code from google.golang.org/grpc/status.
// For connect: connect.CodeOf from connectrpc.com/connect.
func (t Transport) CodeOfIdent() protogen.GoIdent {
	if t.IsConnect() {
		return protogen.GoIdent{
			GoName:       "CodeOf",
			GoImportPath: connectImportPath,
		}
	}
	return protogen.GoIdent{
		GoName:       "Code",
		GoImportPath: "google.golang.org/grpc/status",
	}
}

// NewRequestIdent returns the GoIdent for connect.NewRequest.
// Only meaningful for standard connect mode.
func (t Transport) NewRequestIdent() protogen.GoIdent {
	return protogen.GoIdent{
		GoName:       "NewRequest",
		GoImportPath: connectImportPath,
	}
}

// OutputPackage returns the Go import path and package name for the generated test file.
// For gRPC, it returns the proto package as-is.
// For connect, it returns the connect sub-package (e.g., foov1connect) to avoid import cycles.
func (t Transport) OutputPackage(
	goImportPath protogen.GoImportPath,
	goPackageName protogen.GoPackageName,
) (protogen.GoImportPath, protogen.GoPackageName) {
	if t.IsConnect() {
		connectPkgName := protogen.GoPackageName(string(goPackageName) + "connect")
		connectPath := protogen.GoImportPath(string(goImportPath) + "/" + string(connectPkgName))
		return connectPath, connectPkgName
	}
	return goImportPath, goPackageName
}
