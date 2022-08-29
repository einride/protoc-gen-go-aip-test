package create

import (
	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
)

// Suite of Create tests.
//nolint: gochecknoglobals
var Suite = suite.Suite{
	Name: "Create",
	Tests: []suite.Test{
		parentMissing,
		parentInvalid,
		createTime,
		persisted,
		userSettableID,
		alreadyExists,
		requiredFields,
		resourceReferences,
	},
}
