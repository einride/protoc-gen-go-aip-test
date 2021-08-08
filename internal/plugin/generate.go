package plugin

import (
	"path/filepath"

	"google.golang.org/protobuf/compiler/protogen"
)

const (
	pkgNameSuffix = "test"
	pkgDir        = "testing"
)

func Generate(plugin *protogen.Plugin) error {
	for _, file := range plugin.Files {
		if len(file.Services) == 0 {
			continue
		}
		f := plugin.NewGeneratedFile(filePath(file), goImportPath(file))
		writeHeader(file, f)
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
}
