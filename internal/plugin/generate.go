package plugin

import (
	"fmt"
	"path/filepath"

	"github.com/einride/protoc-gen-go-aiptest/internal/xrange"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const (
	pkgNameSuffix = "test"
	pkgDir        = "testing"
)

func Generate(plugin *protogen.Plugin) error {
	pkgResources := findResourcesPerPackage(plugin)
	for _, file := range plugin.Files {
		if len(file.Services) == 0 || !file.Generate {
			continue
		}
		f := plugin.NewGeneratedFile(filePath(file), goImportPath(file))
		writeHeader(file, f)
		f.Skip()

		for _, service := range file.Services {
			resources := pkgResources[file.Desc.Package()]
			if len(resources) == 0 {
				// no resources in this package.
				continue
			}
			serviceResources := make([]resource, 0, len(resources))
			for _, r := range resources {
				if hasAnyStandardMethodFor(service.Desc, r.descriptor) {
					serviceResources = append(serviceResources, r)
				}
			}
			if len(serviceResources) == 0 {
				continue
			}
			ms := make([]*protogen.Message, 0, len(serviceResources))
			rs := make([]*annotations.ResourceDescriptor, 0, len(serviceResources))
			for _, serviceResource := range serviceResources {
				rs = append(rs, serviceResource.descriptor)
				m, err := protogenMessage(plugin, serviceResource.message.FullName())
				if err != nil {
					return err
				}
				ms = append(ms, m)
			}
			generator := serviceGenerator{
				service:   service,
				resources: rs,
				messages:  ms,
			}
			if err := generator.Generate(f); err != nil {
				return err
			}
			f.Unskip()
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
	f.P("// Code generated by protoc-gen-go-aiptest. DO NOT EDIT.")
	f.P()
	f.P("package ", file.GoPackageName+pkgNameSuffix)
	f.P()
}

func protogenMessage(plugin *protogen.Plugin, name protoreflect.FullName) (*protogen.Message, error) {
	for _, file := range plugin.Files {
		for _, message := range file.Messages {
			if message.Desc.FullName() == name {
				return message, nil
			}
		}
	}
	return nil, fmt.Errorf("no message named '%s' in plugin", name)
}

type resource struct {
	message    protoreflect.MessageDescriptor
	descriptor *annotations.ResourceDescriptor
}

func findResourcesPerPackage(plugin *protogen.Plugin) map[protoreflect.FullName][]resource {
	resources := make(map[protoreflect.FullName][]resource)
	for _, file := range plugin.Files {
		pkg := file.Desc.Package()
		xrange.RangeResourceDescriptors(
			file.Desc,
			func(m protoreflect.MessageDescriptor, r *annotations.ResourceDescriptor) {
				// ignore forwarded resource descriptors
				if m == nil {
					return
				}
				resources[pkg] = append(resources[pkg], resource{
					message:    m,
					descriptor: r,
				})
			},
		)
	}
	return resources
}
