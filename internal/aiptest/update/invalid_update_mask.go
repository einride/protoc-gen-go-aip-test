package update

import (
	"strconv"

	"github.com/einride/protoc-gen-go-aip-test/internal/ident"
	"github.com/einride/protoc-gen-go-aip-test/internal/onlyif"
	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
	"github.com/einride/protoc-gen-go-aip-test/internal/util"
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/compiler/protogen"
)

//nolint: gochecknoglobals
var invalidUpdateMask = suite.Test{
	Name: "invalid update mask",
	Doc: []string{
		"The method should fail with InvalidArgument if the update_mask is invalid.",
	},
	OnlyIf: suite.OnlyIfs(
		onlyif.HasMethod(aipreflect.MethodTypeUpdate),
		onlyif.HasUpdateMask,
	),
	Generate: func(f *protogen.GeneratedFile, scope suite.Scope) error {
		updateMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeUpdate)
		util.MethodUpdate{
			Resource:   scope.Resource,
			Method:     updateMethod,
			Msg:        "created",
			UpdateMask: []string{strconv.Quote("invalid_field_xyz")},
		}.Generate(f, "_", "err", ":=")
		f.P(ident.AssertEqual, "(t, ", ident.Codes(codes.InvalidArgument), ",", ident.StatusCode, "(err), err)")
		return nil
	},
}
