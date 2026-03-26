package main

import (
	"fmt"

	"github.com/einride/protoc-gen-go-aip-test/internal/plugin"
	"github.com/einride/protoc-gen-go-aip-test/internal/transport"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	var transportFlag string
	protogen.Options{
		ParamFunc: func(name, value string) error {
			if name == "transport" {
				transportFlag = value
				return nil
			}
			return fmt.Errorf("unknown parameter %q", name)
		},
	}.Run(func(p *protogen.Plugin) error {
		t, err := transport.Parse(transportFlag)
		if err != nil {
			return err
		}
		return plugin.Generate(p, t)
	})
}
