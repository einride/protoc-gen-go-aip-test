package deletion

import (
	"github.com/einride/protoc-gen-go-aip-test/internal/ident"
	"github.com/einride/protoc-gen-go-aip-test/internal/onlyif"
	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
	"github.com/einride/protoc-gen-go-aip-test/internal/util"
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/protobuf/compiler/protogen"
)

//nolint:gochecknoglobals
var softDeletedDeleteTime = suite.Test{
	Name: "soft-deleted delete_time",
	Doc: []string{
		"A soft-deleted resource should have delete_time assigned.",
	},
	OnlyIf: suite.OnlyIfs(
		onlyif.HasMethod(aipreflect.MethodTypeDelete),
		onlyif.HasField("delete_time"),
		onlyif.ReturnsNotEmpty(aipreflect.MethodTypeDelete),
		onlyif.MethodNotLRO(aipreflect.MethodTypeDelete),
	),
	Generate: func(f *protogen.GeneratedFile, scope suite.Scope) error {
		deleteMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeDelete)

		if util.HasParent(scope.Resource) {
			f.P("parent := ", ident.FixtureNextParent, "(t, false)")
			f.P("created := fx.create(t, parent)")
		} else {
			f.P("created := fx.create(t)")
		}
		f.P("beforeDelete := ", ident.TimeNow, "()")
		util.MethodDelete{
			Resource:    scope.Resource,
			Method:      deleteMethod,
			ResourceVar: "created",
		}.Generate(f, "deleted", "err", ":=")
		f.P(ident.AssertNilError, "(t, err)")
		f.P(ident.AssertCheck, "(t, deleted.DeleteTime.AsTime().After(beforeDelete))")
		return nil
	},
}
