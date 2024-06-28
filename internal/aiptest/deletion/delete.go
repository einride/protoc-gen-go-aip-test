package deletion

import (
	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
)

// Suite of Delete tests.
//
//nolint:gochecknoglobals
var Suite = suite.Suite{
	Name: "Delete",
	Tests: []suite.Test{
		missingName,
		invalidName,
		exists,
		notFound,
		alreadyDeleted,
		wildcardName,
		etagMismatch,
		etagCurrent,
	},
}
