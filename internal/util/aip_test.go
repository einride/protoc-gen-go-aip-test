package util

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test_isSingleton(t *testing.T) {
	for _, tt := range []struct {
		name         string
		resourceName string
		isSingleton  bool
	}{
		{
			name:         "empty",
			resourceName: "",
			isSingleton:  false,
		},
		{
			name:         "top-level singleton",
			resourceName: "settings",
			isSingleton:  true,
		},
		{
			name:         "resource",
			resourceName: "shippers/{shipper}",
			isSingleton:  false,
		},
		{
			name:         "singleton resource",
			resourceName: "shippers/{shipper}/settings",
			isSingleton:  true,
		},
		{
			name:         "resource with child",
			resourceName: "shippers/{shipper}/shipments/{shipment}",
			isSingleton:  false,
		},
		{
			name:         "singleton resource with child",
			resourceName: "shippers/{shipper}/shipments/{shipment}/settings",
			isSingleton:  true,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.isSingleton, isSingleton(tt.resourceName))
		})
	}
}
