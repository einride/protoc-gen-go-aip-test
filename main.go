package main

import (
	"flag"

	"github.com/einride/protoc-gen-go-aip-test/internal/plugin"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	flagSet := flag.NewFlagSet("config", flag.ContinueOnError)
	var config plugin.Config
	config.AddToFlagSet(flagSet)
	opts := protogen.Options{ParamFunc: flagSet.Set}
	opts.Run(func(gen *protogen.Plugin) error {
		return plugin.GenerateWithConfig(gen, config)
	})
}
