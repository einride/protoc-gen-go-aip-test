package suite

import (
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
)

type Scope struct {
	Service  *protogen.Service
	Resource *annotations.ResourceDescriptor
	Message  *protogen.Message
}

// Suite contains a suite of tests for a method.
type Suite struct {
	Name       string
	Tests      []Test
	TestGroups []TestGroup
}

func (m Suite) Enabled(scope Scope) bool {
	for _, t := range m.Tests {
		if t.Enabled(scope) {
			return true
		}
	}
	for _, tg := range m.TestGroups {
		if tg.Enabled(scope) {
			return true
		}
	}
	return false
}

// Test is one test in the generated code.
type Test struct {
	Name     string
	Doc      []string
	OnlyIf   func(scope Scope) bool
	Generate func(f *protogen.GeneratedFile, scope Scope) error
}

func (t Test) Enabled(scope Scope) bool {
	return t.OnlyIf(scope)
}

// TestGroup contains multiple tests in the generated code
// that share some setup code.
type TestGroup struct {
	OnlyIf         func(scope Scope) bool
	GenerateBefore func(f *protogen.GeneratedFile, scope Scope) error
	Tests          []Test
}

func (tg TestGroup) Enabled(scope Scope) bool {
	if tg.OnlyIf != nil && !tg.OnlyIf(scope) {
		return false
	}
	for _, t := range tg.Tests {
		if t.Enabled(scope) {
			return true
		}
	}
	return false
}
