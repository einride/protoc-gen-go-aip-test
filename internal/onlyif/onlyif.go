package onlyif

import (
	"fmt"

	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
	"github.com/einride/protoc-gen-go-aip-test/internal/util"
	"go.einride.tech/aip/reflect/aipreflect"
)

var _ suite.OnlyIf = onlyIf{}

type onlyIf struct {
	f   func(scope suite.Scope) bool
	doc string
}

func (o onlyIf) Check(scope suite.Scope) bool {
	return o.f(scope)
}

func (o onlyIf) String() string {
	return o.doc
}

func HasMethod(methodType aipreflect.MethodType) suite.OnlyIf {
	return onlyIf{
		f: func(scope suite.Scope) bool {
			_, ok := util.StandardMethod(scope.Service, scope.Resource, methodType)
			return ok
		},
		doc: fmt.Sprintf("has %s method", methodType),
	}
}

func MethodNotLRO(methodType aipreflect.MethodType) suite.OnlyIf {
	return onlyIf{
		f: func(scope suite.Scope) bool {
			method, ok := util.StandardMethod(scope.Service, scope.Resource, methodType)
			return ok && !util.ReturnsLRO(method.Desc)
		},
		doc: fmt.Sprintf("%s method does not return long-running operation", methodType),
	}
}

var HasUserSettableID = onlyIf{
	f: func(scope suite.Scope) bool {
		method, ok := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeCreate)
		return ok && util.HasUserSettableIDField(scope.Resource, method.Input.Desc)
	},
	doc: "has user settable ID",
}

var HasUpdateMask = onlyIf{
	f: func(scope suite.Scope) bool {
		method, ok := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeUpdate)
		return ok && util.HasUpdateMask(method.Desc)
	},
	doc: "Update method has update_mask",
}

var HasParent = onlyIf{
	f: func(scope suite.Scope) bool {
		return util.HasParent(scope.Resource)
	},
	doc: "resource has a parent",
}

var HasRequiredFields = onlyIf{
	f: func(scope suite.Scope) bool {
		return util.HasRequiredFields(scope.Message.Desc)
	},
	doc: "resource has any required fields",
}

var HasMutableResourceReferences = onlyIf{
	f: func(scope suite.Scope) bool {
		return util.HasMutableResourceReferences(scope.Message.Desc)
	},
	doc: "resource has any mutable resource references",
}

func BatchMethodNotAlternative(methodType aipreflect.MethodType) suite.OnlyIf {
	return onlyIf{
		f: func(scope suite.Scope) bool {
			method, ok := util.StandardMethod(scope.Service, scope.Resource, methodType)
			if !ok {
				return false
			}
			return !util.IsAlternativeBatch(method.Desc)
		},
		doc: "is not alternative batch request message",
	}
}
