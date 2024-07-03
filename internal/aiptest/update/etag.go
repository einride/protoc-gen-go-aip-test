package update

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
		onlyif.HasMethod(aipreflect.MethodTypeUpdate),
		onlyif.HasRequestEtag(aipreflect.MethodTypeUpdate),
		onlyif.HasField("etag"),
	),
	Generate: func(f *protogen.GeneratedFile, scope suite.Scope) error {
		if util.HasParent(scope.Resource) {
			f.P("parent := ", ident.FixtureNextParent, "(t, false)")
			f.P("created := fx.create(t, parent)")
		} else {
			f.P("created := fx.create(t)")
		}
		updateMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeUpdate)
		util.MethodUpdate{
			Resource: scope.Resource,
			Method:   updateMethod,
			Msg:      "created",
			Etag:     util.EtagLiteral("99999"),
		}.Generate(f, "_", "err", ":=")
		f.P(ident.AssertEqual, "(t, ", ident.Codes(codes.Aborted), ",", ident.StatusCode, "(err), err)")
		return nil
	},
}

//nolint:gochecknoglobals
var etagUpdated = suite.Test{
	Name: "etag updated",
	Doc: []string{
		"Field etag should have a new value when the resource is successfully updated.",
	},
	OnlyIf: suite.OnlyIfs(
		onlyif.HasMethod(aipreflect.MethodTypeUpdate),
		onlyif.HasRequestEtag(aipreflect.MethodTypeUpdate),
		onlyif.HasField("etag"),
	),
	Generate: func(f *protogen.GeneratedFile, scope suite.Scope) error {
		if util.HasParent(scope.Resource) {
			f.P("parent := ", ident.FixtureNextParent, "(t, false)")
			f.P("created := fx.create(t, parent)")
		} else {
			f.P("created := fx.create(t)")
		}
		updateMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeUpdate)
		util.MethodUpdate{
			Resource: scope.Resource,
			Method:   updateMethod,
			Parent:   "parent",
			Name:     "created.Name",
			Etag:     "created.Etag",
		}.Generate(f, "updated", "err", ":=")
		f.P(ident.AssertNilError, "(t, err)")
		f.P(ident.AssertCheck, "(t, updated.Etag != created.Etag)")
		return nil
	},
}
