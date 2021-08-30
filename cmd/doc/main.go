package main

import (
	"fmt"
	"strings"

	"github.com/einride/protoc-gen-go-aip-test/internal/aiptest"
)

func main() {
	for _, suite := range aiptest.Suites {
		fmt.Println("###", suite.Name)
		fmt.Println("| Name | Description |")
		fmt.Println("| ---- | ----------- |")
		for _, test := range suite.Tests {
			fmt.Println("|", test.Name, "|", strings.Join(test.Doc, " "), "|")
		}
		for _, group := range suite.TestGroups {
			for _, test := range group.Tests {
				fmt.Println("|", test.Name, "|", strings.Join(test.Doc, " "), "|")
			}
		}
	}
}
