package main

import (
	"fmt"
	"strings"

	"github.com/einride/protoc-gen-go-aip-test/internal/aiptest"
	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
)

func main() {
	for _, s := range aiptest.Suites {
		fmt.Println("###", s.Name)
		fmt.Println("| Name | Description | Only if |")
		fmt.Println("| ---- | ----------- | ------- |")
		for _, test := range s.Tests {
			printTestRow(test.Name, test.Doc, test.OnlyIf)
		}
		for _, group := range s.TestGroups {
			for _, test := range group.Tests {
				testOnlyIf := test.OnlyIf
				if group.OnlyIf != nil {
					testOnlyIf = suite.OnlyIfs(group.OnlyIf, test.OnlyIf)
				}
				printTestRow(test.Name, test.Doc, testOnlyIf)
			}
		}
	}
}

func printTestRow(name string, doc []string, onlyif suite.OnlyIf) {
	fmt.Println(
		"|",
		name,
		"|",
		strings.Join(doc, " "),
		"|",
		"Generated only if all are true:",
		formatOnlyIfMarkdownList(onlyif),
		"|",
	)
}

func formatOnlyIfMarkdownList(onlyIf suite.OnlyIf) string {
	if composed, ok := onlyIf.(suite.ComposedOnlyIf); ok {
		onlyIfs := composed.Flat()
		onlyIfsStr := make([]string, 0, len(onlyIfs))
		for _, o := range onlyIfs {
			onlyIfsStr = append(onlyIfsStr, "<li>"+o.String()+"</li>")
		}
		return "<ul>" + strings.Join(onlyIfsStr, "") + "</ul>"
	}
	return onlyIf.String()
}
