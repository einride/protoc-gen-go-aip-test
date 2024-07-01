package deletion

import (
	"github.com/einride/protoc-gen-go-aip-test/internal/ident"
	"github.com/einride/protoc-gen-go-aip-test/internal/onlyif"
	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
	"github.com/einride/protoc-gen-go-aip-test/internal/util"
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/compiler/protogen"
)

//nolint:gochecknoglobals
var etagMismatch = suite.Test{
	Name: "etag mismatch",
	Doc: []string{
		"Method should fail with Aborted if the supplied etag doesnt match the current etag value.",
	},
	OnlyIf: suite.OnlyIfs(
		onlyif.HasMethod(aipreflect.MethodTypeDelete),
		onlyif.HasRequestEtag(aipreflect.MethodTypeDelete),
		onlyif.HasField("etag"),
	),
	Generate: func(f *protogen.GeneratedFile, scope suite.Scope) error {
		if util.HasParent(scope.Resource) {
			f.P("parent := ", ident.FixtureNextParent, "(t, false)")
			f.P("created := fx.create(t, parent)")
		} else {
			f.P("created := fx.create(t)")
		}
		deleteMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeDelete)
		util.MethodDelete{
			Resource:         scope.Resource,
			Method:           deleteMethod,
			ResourceVar:      "created",
			UserProvidedEtag: util.EtagLiteral("99999"),
		}.Generate(f, "_", "err", ":=")
		f.P(ident.AssertEqual, "(t, ", ident.Codes(codes.Aborted), ",", ident.StatusCode, "(err), err)")
		return nil
	},
}

//nolint:gochecknoglobals
var etagCurrent = suite.Test{
	Name: "current etag",
	Doc: []string{
		"Deletion with the current etag supplied should succeed.",
	},
	OnlyIf: suite.OnlyIfs(
		onlyif.HasMethod(aipreflect.MethodTypeDelete),
		onlyif.HasRequestEtag(aipreflect.MethodTypeDelete),
		onlyif.HasField("etag"),
	),
	Generate: func(f *protogen.GeneratedFile, scope suite.Scope) error {
		if util.HasParent(scope.Resource) {
			f.P("parent := ", ident.FixtureNextParent, "(t, false)")
			f.P("created := fx.create(t, parent)")
		} else {
			f.P("created := fx.create(t)")
		}
		deleteMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeDelete)
		util.MethodDelete{
			Resource:         scope.Resource,
			Method:           deleteMethod,
			ResourceVar:      "created",
			UserProvidedEtag: "created.Etag",
		}.Generate(f, "_", "err", ":=")
		f.P(ident.AssertNilError, "(t, err)")
		return nil
	},
}
