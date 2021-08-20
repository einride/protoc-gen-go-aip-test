package get

import (
	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
)

var Suite = suite.Suite{
	Name: "Get",
	Tests: []suite.Test{
		missingName,
		invalidName,
		exists,
		notFound,
		// TODO: add test for supplying wildcard as name
	},
}
