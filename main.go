package main

import (
	"github.com/einride/protoc-gen-go-aiptest/internal/plugin"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	protogen.Options{}.Run(plugin.Generate)
}
