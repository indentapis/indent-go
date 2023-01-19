package main

import (
	"flag"

	"google.golang.org/protobuf/compiler/protogen"

	"go.indent.com/indent-go/api/indent/log/pkg/plugin"
)

func main() {
	var (
		flags        flag.FlagSet
		conservative = flags.Bool(
			"conservative",
			true,
			"only define ObjectMarshaler for messages with logged fields",
		)
	)

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			plugin.GenerateFile(gen, f, *conservative)
		}
		return nil
	})
}
