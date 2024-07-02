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
var alreadyDeleted = suite.Test{
	Name: "already deleted",
	Doc: []string{
		"Method should fail with NotFound if the resource was already deleted. This also applies to soft-deletion.",
	},

	OnlyIf: suite.OnlyIfs(
		onlyif.HasMethod(aipreflect.MethodTypeDelete),
	),
	Generate: func(f *protogen.GeneratedFile, scope suite.Scope) error {
		deleteMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeDelete)
		if util.HasParent(scope.Resource) {
			f.P("parent := ", ident.FixtureNextParent, "(t, false)")
			f.P("created := fx.create(t, parent)")
		} else {
			f.P("created := fx.create(t)")
		}
		responseVariable := "_"
		deletedEtag := ""
		if util.HasEtagField(deleteMethod.Input.Desc) && util.HasEtagField(deleteMethod.Output.Desc) {
			// Second call to delete we need to define response variable which can be used to extract the etag for
			// the next request.
			// Only create variable if both request and response contain an etag field.
			responseVariable = "deleted"
			deletedEtag = "deleted.Etag"
		}
		util.MethodDelete{
			Resource:    scope.Resource,
			Method:      deleteMethod,
			ResourceVar: "created",
		}.Generate(f, responseVariable, "err", ":=")
		f.P(ident.AssertNilError, "(t, err)")
		util.MethodDelete{
			Resource:         scope.Resource,
			Method:           deleteMethod,
			ResourceVar:      "created",
			UserProvidedEtag: deletedEtag,
		}.Generate(f, "_", "err", "=")
		f.P(ident.AssertEqual, "(t, ", ident.Codes(codes.NotFound), ",", ident.StatusCode, "(err), err)")
		return nil
	},
}
