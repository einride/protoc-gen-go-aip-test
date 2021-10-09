package suite

import (
	"strings"
)

type OnlyIf interface {
	Check(scope Scope) bool
	String() string
}

type ComposedOnlyIf interface {
	Flat() []OnlyIf
}

func OnlyIfs(onlyIfs ...OnlyIf) OnlyIf {
	return onlyIf{
		children: onlyIfs,
	}
}

var (
	_ OnlyIf         = onlyIf{}
	_ ComposedOnlyIf = onlyIf{}
)

type onlyIf struct {
	children []OnlyIf
}

func (o onlyIf) Flat() []OnlyIf {
	flat := make([]OnlyIf, 0, len(o.children))
	for _, child := range o.children {
		if composedChild, ok := child.(ComposedOnlyIf); ok {
			flat = append(flat, composedChild.Flat()...)
		} else {
			flat = append(flat, child)
		}
	}
	return flat
}

func (o onlyIf) Check(scope Scope) bool {
	for _, child := range o.children {
		if !child.Check(scope) {
			return false
		}
	}
	return true
}

func (o onlyIf) String() string {
	docs := make([]string, 0, len(o.children))
	for _, child := range o.children {
		docs = append(docs, child.String())
	}
	return strings.Join(docs, " and ")
}
