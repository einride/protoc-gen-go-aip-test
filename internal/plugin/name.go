package plugin

import (
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/genproto/googleapis/api/annotations"
)

func resourceType(r *annotations.ResourceDescriptor) string {
	return aipreflect.ResourceType(r.GetType()).Type()
}
