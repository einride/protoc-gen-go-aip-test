// Code generated by protoc-gen-go-aip-test. DO NOT EDIT.

package pubsubpb

import (
	testing "testing"
)

// ServiceConfigProviders embeds providers for all services.
type ServiceConfigProviders interface {
	SchemaServiceTestSuiteConfigProvider
	PublisherTestSuiteConfigProvider
	SubscriberTestSuiteConfigProvider
}

// TestServices is the main entrypoint for starting the AIP tests for all services.
func TestServices(t *testing.T, s ServiceConfigProviders) {
	TestSchemaService(t, s)
	TestPublisher(t, s)
	TestSubscriber(t, s)
}
