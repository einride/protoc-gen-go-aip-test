package update

import (
	"strconv"

	"github.com/einride/protoc-gen-go-aip-test/internal/ident"
	"github.com/einride/protoc-gen-go-aip-test/internal/onlyif"
	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
	"github.com/einride/protoc-gen-go-aip-test/internal/util"
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/protobuf/compiler/protogen"
)

//nolint:gochecknoglobals
var preserveCreateTime = suite.Test{
	Name: "preserve create_time",
	Doc: []string{
		"The field create_time should be preserved when a '*'-update mask is used.",
	},

	OnlyIf: suite.OnlyIfs(
		onlyif.HasMethod(aipreflect.MethodTypeUpdate),
		onlyif.MethodNotLRO(aipreflect.MethodTypeUpdate),
		onlyif.HasField("create_time"),
		onlyif.HasRequiredFields,
	),
	Generate: func(f *protogen.GeneratedFile, scope suite.Scope) error {
		updateMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeUpdate)

		if util.HasParent(scope.Resource) {
			f.P("parent := ", ident.FixtureNextParent, "(t, false)")
			f.P("created := fx.create(t, parent)")
		} else {
			f.P("created := fx.create(t)")
		}
		f.P("originalCreateTime := created.CreateTime")
		util.MethodUpdate{
			Resource:   scope.Resource,
			Method:     updateMethod,
			Msg:        "created",
			UpdateMask: []string{strconv.Quote("*")},
		}.Generate(f, "updated", "err", ":=")
		f.P(ident.AssertNilError, "(t, err)")
		f.P(ident.AssertDeepEqual, "(t, originalCreateTime, updated.CreateTime,", ident.ProtocmpTransform, "())")
		return nil
	},
}
