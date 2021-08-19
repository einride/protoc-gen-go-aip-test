package plugin

import (
	"strconv"

	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protopath"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (r *resourceGenerator) updateTestCase() testCase {
	updateMethod, ok := r.standardMethod(aipreflect.MethodTypeUpdate)
	if !ok {
		return disabledTestCase()
	}
	createMethod, ok := r.standardMethod(aipreflect.MethodTypeCreate)
	if !ok {
		return disabledTestCase()
	}
	// TODO: support LROs for create.
	if returnsLRO(createMethod.Desc) {
		return disabledTestCase()
	}
	getMethod, hasGet := r.standardMethod(aipreflect.MethodTypeGet)
	isLRO := returnsLRO(updateMethod.Desc)

	return newTestCase("Update", func(f *protogen.GeneratedFile) {
		testingT := f.QualifiedGoIdent(protogen.GoIdent{GoName: "T", GoImportPath: "testing"})
		assertEqual := f.QualifiedGoIdent(protogen.GoIdent{
			GoName:       "Equal",
			GoImportPath: "gotest.tools/v3/assert",
		})
		assertCheck := f.QualifiedGoIdent(protogen.GoIdent{
			GoName:       "Check",
			GoImportPath: "gotest.tools/v3/assert",
		})
		assertDeepEqual := f.QualifiedGoIdent(protogen.GoIdent{
			GoName:       "DeepEqual",
			GoImportPath: "gotest.tools/v3/assert",
		})
		assertNilError := f.QualifiedGoIdent(protogen.GoIdent{
			GoName:       "NilError",
			GoImportPath: "gotest.tools/v3/assert",
		})
		protocmpTransform := f.QualifiedGoIdent(protogen.GoIdent{
			GoName:       "Transform",
			GoImportPath: "google.golang.org/protobuf/testing/protocmp",
		})
		protoClone := f.QualifiedGoIdent(protogen.GoIdent{
			GoName:       "Clone",
			GoImportPath: "google.golang.org/protobuf/proto",
		})
		statusCode := f.QualifiedGoIdent(protogen.GoIdent{
			GoName:       "Code",
			GoImportPath: "google.golang.org/grpc/status",
		})
		codesInvalidArgument := f.QualifiedGoIdent(protogen.GoIdent{
			GoName:       "InvalidArgument",
			GoImportPath: "google.golang.org/grpc/codes",
		})
		codesNotFound := f.QualifiedGoIdent(protogen.GoIdent{
			GoName:       "NotFound",
			GoImportPath: "google.golang.org/grpc/codes",
		})

		f.P("// Standard methods: Update")
		f.P("// https://google.aip.dev/134")

		if hasParent(r.resource) {
			f.P()
			f.P("parent := fx.nextParent(t, false)")
		}
		methodCreate{
			resource: r.resource,
			method:   createMethod,
			parent:   "parent",
		}.Generate(f, "created00", "err", ":=")
		f.P(assertNilError, "(t, err)")

		f.P()
		f.P("// Method should fail with InvalidArgument if no name is provided.")
		f.P("t.Run(\"missing name\", func(t *", testingT, ") {")
		f.P("fx.maybeSkip(t)")
		methodUpdate{
			resource: r.resource,
			method:   updateMethod,

			parent: "parent",
			name:   strconv.Quote(""),
		}.Generate(f, "_", "err", ":=")
		f.P(assertEqual, "(t, ", codesInvalidArgument, ",", statusCode, "(err), err)")
		f.P("})")

		f.P()
		f.P("// Method should fail with InvalidArgument is provided name is not valid.")
		f.P("t.Run(\"invalid name\", func(t *", testingT, ") {")
		f.P("fx.maybeSkip(t)")
		methodUpdate{
			resource: r.resource,
			method:   updateMethod,
			parent:   "parent",
			name:     strconv.Quote("invalid resource name"),
		}.Generate(f, "_", "err", ":=")
		f.P(assertEqual, "(t, ", codesInvalidArgument, ",", statusCode, "(err), err)")
		f.P("})")

		if !isLRO {
			f.P()
			f.P("// Field update_time should be updated when the resource is updated.")
			f.P("t.Run(\"update time\", func(t *", testingT, ") {")
			f.P("fx.maybeSkip(t)")
			methodCreate{
				resource: r.resource,
				method:   createMethod,
				parent:   "parent",
			}.Generate(f, "initial", "err", ":=")
			f.P(assertNilError, "(t, err)")
			methodUpdate{
				resource: r.resource,
				method:   updateMethod,
				msg:      "initial",
			}.Generate(f, "updated", "err", ":=")
			f.P(assertNilError, "(t, err)")
			f.P(assertCheck, "(t, updated.UpdateTime.AsTime().After(initial.UpdateTime.AsTime()))")
			f.P("})")
		}

		f.P()
		f.P("// Method should fail with NotFound if the resource does not exist.")
		f.P("t.Run(\"not found\", func(t *", testingT, ") {")
		f.P("fx.maybeSkip(t)")
		methodUpdate{
			resource: r.resource,
			method:   updateMethod,
			parent:   "parent",
			// appending to the resource name ensures it is valid
			name: "created00.Name + \"notfound\"",
		}.Generate(f, "_", "err", ":=")
		f.P(assertEqual, "(t, ", codesNotFound, ",", statusCode, "(err), err)")
		f.P("})")

		if hasGet && !isLRO {
			f.P()
			f.P("// The updated resource should be persisted and reachable with Get.")
			f.P("t.Run(\"persisted\", func(t *", testingT, ") {")
			f.P("fx.maybeSkip(t)")
			methodCreate{
				resource: r.resource,
				method:   createMethod,
				parent:   "parent",
			}.Generate(f, "initial", "err", ":=")
			f.P(assertNilError, "(t, err)")
			methodUpdate{
				resource: r.resource,
				method:   updateMethod,
				msg:      "initial",
			}.Generate(f, "updated", "err", ":=")
			methodGet{
				resource: r.resource,
				method:   getMethod,
				name:     "updated.Name",
			}.Generate(f, "persisted", "err", ":=")
			f.P(assertNilError, "(t, err)")
			f.P(assertDeepEqual, "(t, updated, persisted, ", protocmpTransform, "())")
			f.P("})")
		}

		if hasUpdateMask(updateMethod.Desc) {
			f.P()
			f.P("// The method should fail with InvalidArgument if the update_mask is invalid.")
			f.P("t.Run(\"invalid update mask\", func(t *", testingT, ") {")
			f.P("fx.maybeSkip(t)")
			methodUpdate{
				resource: r.resource,
				method:   updateMethod,
				// appending to the resource name ensures it is valid
				msg:        "created00",
				updateMask: []string{strconv.Quote("invalid_field_xyz")},
			}.Generate(f, "_", "err", ":=")
			f.P(assertEqual, "(t, ", codesInvalidArgument, ",", statusCode, "(err), err)")
			f.P("})")
		}

		if hasUpdateMask(updateMethod.Desc) && hasRequiredFields(r.message.Desc) {
			f.P()
			f.P("// Method should fail with InvalidArgument if any required field is missing")
			f.P("// when called with '*' update_mask.")
			f.P("t.Run(\"required fields\", func(t *", testingT, ") {")
			f.P("fx.maybeSkip(t)")
			rangeRequiredFields(r.message.Desc, func(p protopath.Path, field protoreflect.FieldDescriptor) {
				// strip root step
				p = p[1:]
				containerPath := p[:len(p)-1]
				fieldPath := p[len(p)-1]
				isTopLevel := len(containerPath) == 0

				f.P("t.Run(", strconv.Quote(p.String()), ", func(t *", testingT, ") {")
				f.P("fx.maybeSkip(t)")
				f.P("msg := ", protoClone, "(created00).(*", r.message.GoIdent, ")")
				if isTopLevel {
					f.P("container := msg")
				} else {
					f.P("container := msg.", chainedGet(containerPath))
				}
				f.P("if container == nil {")
				f.P("t.Skip(\"not reachable\")")
				f.P("}")
				fieldName := string(fieldPath.FieldDescriptor().Name())
				f.P("fd := container.ProtoReflect().Descriptor().Fields().ByName(", strconv.Quote(fieldName), ")")
				f.P("container.ProtoReflect().Clear(fd)")
				m := methodUpdate{
					resource:   r.resource,
					method:     createMethod,
					msg:        "msg",
					updateMask: []string{strconv.Quote("*")},
				}
				m.Generate(f, "_", "err", ":=")
				f.P(assertEqual, "(t, ", codesInvalidArgument, ", ", statusCode, "(err), err)")
				f.P("})")
			})
			f.P("})")
		}

		// TODO: add test for supplying wildcard as name
		// TODO: add test for etags

		f.P("_ = ", codesNotFound)
		f.P("_ = ", protocmpTransform)
		f.P("_ = ", protoClone)
	})
}
