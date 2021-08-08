package plugin

import (
	"fmt"
	"path/filepath"
	"sort"

	"go.einride.tech/aip/reflect/aipreflect"
	"go.einride.tech/aip/reflect/aipregistry"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

const (
	pkgNameSuffix = "test"
	pkgDir        = "testing"
)

func Generate(plugin *protogen.Plugin) error {
	protoRegistry, err := protoRegistryFromPlugin(plugin)
	if err != nil {
		return err
	}
	aipRegistry, err := aipregistry.NewResources(protoRegistry)
	if err != nil {
		return fmt.Errorf("initialize AIP registry: %w", err)
	}

	for _, file := range plugin.Files {
		if len(file.Services) == 0 || !file.Generate {
			continue
		}
		f := plugin.NewGeneratedFile(filePath(file), goImportPath(file))
		writeHeader(file, f)

		for _, service := range file.Services {
			resources := findServiceResources(aipRegistry, service.Desc.FullName())
			if len(resources) == 0 {
				continue
			}
			messages, err := findResourceMessages(protoRegistry, resources)
			if err != nil {
				return err
			}
			generator := serviceGenerator{
				service:   service,
				resources: resources,
				messages:  messages,
			}
			if err := generator.Generate(f); err != nil {
				return err
			}
		}
	}
	return nil
}

func filePath(file *protogen.File) string {
	dir := filepath.Dir(file.GeneratedFilenamePrefix)
	filePrefix := filepath.Base(file.GeneratedFilenamePrefix)

	return filepath.Join(dir, pkgDir, filePrefix+".go")
}

func goImportPath(file *protogen.File) protogen.GoImportPath {
	return protogen.GoImportPath(filepath.Base(filePath(file)))
}

func writeHeader(file *protogen.File, f *protogen.GeneratedFile) {
	f.P("package ", file.GoPackageName+pkgNameSuffix)
	f.P()
}

func protoRegistryFromPlugin(plugin *protogen.Plugin) (*protoregistry.Files, error) {
	var protoReg protoregistry.Files
	for _, file := range plugin.Files {
		if err := protoReg.RegisterFile(file.Desc); err != nil {
			return nil, fmt.Errorf("register proto file: %w", err)
		}
	}
	return &protoReg, nil
}

func findServiceResources(
	resources *aipregistry.Resources,
	service protoreflect.FullName,
) []*aipreflect.ResourceDescriptor {
	var found []*aipreflect.ResourceDescriptor
	resources.RangeResources(func(descriptor *aipreflect.ResourceDescriptor) bool {
		for _, method := range descriptor.Methods {
			if method.Parent() == service {
				found = append(found, descriptor)
				return true
			}
		}
		return true
	})
	sort.Slice(found, func(i, j int) bool {
		return found[i].Type < found[j].Type
	})
	return found
}

func findResourceMessages(
	registry *protoregistry.Files,
	resources []*aipreflect.ResourceDescriptor,
) ([]protoreflect.MessageDescriptor, error) {
	msgs := make([]protoreflect.MessageDescriptor, 0, len(resources))
	for _, resource := range resources {
		msg, err := registry.FindDescriptorByName(resource.Message)
		if err != nil {
			return nil, fmt.Errorf("find descriptor for resource '%s': %w", resource.Type.Type(), err)
		}
		msgs = append(msgs, msg.(protoreflect.MessageDescriptor))
	}
	return msgs, nil
}
