// Code generated by protoc-gen-go-aip-test. DO NOT EDIT.

package gsuiteaddonspb

import (
	testing "testing"
)

// ServiceConfigProviders embeds providers for all services.
type ServiceConfigProviders interface {
	GSuiteAddOnsTestSuiteConfigProvider
}

// TestServices is the main entrypoint for starting the AIP tests for all services.
func TestServices(t *testing.T, s ServiceConfigProviders) {
	testGSuiteAddOns(t, s)
}
