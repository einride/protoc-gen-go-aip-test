package get

import (
	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
)

// Suite of Get tests.
// nolint: gochecknoglobals
var Suite = suite.Suite{
	Name: "Get",
	Tests: []suite.Test{
		missingName,
		invalidName,
		exists,
		notFound,
		wildcardName,
	},
}
