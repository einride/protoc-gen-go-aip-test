package get

import (
	"github.com/einride/protoc-gen-go-aip-test/internal/ident"
	"github.com/einride/protoc-gen-go-aip-test/internal/onlyif"
	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
	"github.com/einride/protoc-gen-go-aip-test/internal/util"
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/protobuf/compiler/protogen"
)

//nolint:gochecknoglobals
var softDeleted = suite.Test{
	Name: "soft-deleted",
	Doc: []string{
		"A soft-deleted resource should be returned without errors.",
	},
	OnlyIf: suite.OnlyIfs(
		onlyif.HasMethod(aipreflect.MethodTypeGet),
		onlyif.HasMethod(aipreflect.MethodTypeDelete),
		onlyif.HasField("delete_time"),
	),
	Generate: func(f *protogen.GeneratedFile, scope suite.Scope) error {
		getMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeGet)
		deleteMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeDelete)

		if util.HasParent(scope.Resource) {
			f.P("parent := ", ident.FixtureNextParent, "(t, false)")
			f.P("created := fx.create(t, parent)")
		} else {
			f.P("created := fx.create(t)")
		}
		util.MethodDelete{
			Resource:    scope.Resource,
			Method:      deleteMethod,
			ResourceVar: "created",
		}.Generate(f, "deleted", "err", ":=")
		f.P(ident.AssertNilError, "(t, err)")
		util.MethodGet{
			Resource: scope.Resource,
			Method:   getMethod,
			Name:     "created.Name",
		}.Generate(f, "msg", "err", ":=")
		f.P(ident.AssertNilError, "(t, err)")
		if util.ReturnsEmpty(deleteMethod.Desc) {
			// skip asserting if the deleted method returns an Empty response.
			f.P("_ = deleted") // prevent unused variable error.
			f.P("_ = msg")     // prevent unused variable error.
		} else {
			f.P(ident.AssertDeepEqual, "(t, msg, deleted, ", ident.ProtocmpTransform, "())")
		}
		return nil
	},
}
