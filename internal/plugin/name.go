package plugin

import (
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func resourceType(r *annotations.ResourceDescriptor) string {
	return aipreflect.ResourceType(r.GetType()).Type()
}

func serviceTestSuiteName(service protoreflect.ServiceDescriptor) string {
	return string(service.Name()) + "TestSuite"
}

func resourceTestSuiteConfigName(
	service protoreflect.ServiceDescriptor,
	resource *annotations.ResourceDescriptor,
) string {
	return string(service.Name()) + resourceType(resource) + "TestSuiteConfig"
}

func serviceTestConfigSupplierName(service protoreflect.ServiceDescriptor) string {
	return string(service.Name()) + "TestsConfigSupplier"
}
