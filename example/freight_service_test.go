package example

import (
	"context"
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

	// run tests for each resource in the service
	ctx := context.Background()
	suite.TestShipper(ctx, examplefreightv1.ShipperTestSuiteConfig{
		// Create should return a resource which is valid to create, i.e.
		// all required fields set.
		Create: func() *examplefreightv1.Shipper {
			return &examplefreightv1.Shipper{
				DisplayName:    "Example shipper",
				BillingAccount: "billingAccounts/12345",
			}
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
