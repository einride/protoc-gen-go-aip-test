package example

import (
	"context"
	"fmt"
	"sync/atomic"
	"testing"

	examplefreightv1 "github.com/einride/protoc-gen-go-aip-test/proto/gen/einride/example/freight/v1"
)

func Test_FreightService(t *testing.T) {
	t.Parallel()
	t.Skip("this is just an example, the service is not implemented.")
	// setup server before test
	server := examplefreightv1.UnimplementedFreightServiceServer{}
	// setup test suite
	suite := examplefreightv1.FreightServiceTestSuite{
		T:      t,
		Server: server,
	}

	// counter to keep track of unique IDs.
	var idCounter uint64

	// run tests for each resource in the service
	ctx := context.Background()
	suite.TestShipper(ctx, examplefreightv1.FreightServiceShipperTestSuiteConfig{
		// Create should return a resource which is valid to create, i.e.
		// all required fields set.
		Create: func() *examplefreightv1.Shipper {
			return &examplefreightv1.Shipper{
				DisplayName:    "Example shipper",
				BillingAccount: "billingAccounts/12345",
			}
		},
		// IDGenerator should return a valid and unique ID to use in the Create call.
		// If non-nil, this function will be called to set the ID on all Create calls.
		// If the ID field is required, tests will fail if this is nil.
		IDGenerator: func() string {
			id := atomic.AddUint64(&idCounter, 1)
			return fmt.Sprintf("valid-id-%d", id)
		},
		// Update should return a resource which is valid to update, i.e.
		// all required fields set.
		Update: func() *examplefreightv1.Shipper {
			return &examplefreightv1.Shipper{
				DisplayName:    "Updated example shipper",
				BillingAccount: "billingAccounts/54321",
			}
		},
	})
}
