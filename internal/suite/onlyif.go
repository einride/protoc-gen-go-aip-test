package suite

import (
	"strings"
)

type OnlyIf interface {
	Check(scope Scope) bool
	String() string
}

func OnlyIfs(onlyIfs ...OnlyIf) OnlyIf {
	docs := make([]string, 0, len(onlyIfs))
	for _, s := range onlyIfs {
		docs = append(docs, s.String())
	}
	return onlyIf{
		f: func(scope Scope) bool {
			for _, o := range onlyIfs {
				if !o.Check(scope) {
					return false
				}
			}
			return true
		},
		doc: strings.Join(docs, " and "),
	}
}

var _ OnlyIf = onlyIf{}

type onlyIf struct {
	f   func(scope Scope) bool
	doc string
}

func (o onlyIf) Check(scope Scope) bool {
	return o.f(scope)
}

func (o onlyIf) String() string {
	return o.doc
}
