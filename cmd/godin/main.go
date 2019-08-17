package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"gitub.com/go-godin/godin"
	grpcServer "gitub.com/go-godin/godin/module/transport/grpc/server"
)

func main() {
	wd, _ := os.Getwd()

	registry := registerModules()
	app := godin.NewGodin(registry, wd, "examples")

	if err := app.EnsureConfigFile(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// pretend to enable modules, this should be done via the godin cli
	if err := app.AddModule("transport.grpc.server"); err != nil {
		panic(err)
	}

	// parse a protobuf, which should also be passed via the cli
	ctx, err := godin.Parse("/home/lukas/devel/work/protobuf/ticket/ticket/ticket.proto")
	if err != nil {
		panic(err)
	}

	//spew.Dump(ctx)

	// ensure we can write the generated files
	if err := app.EnsureOutputPath(); err != nil {
		fmt.Println(err)
	}

	// generate all enabled modules
	for _, module := range registry.Modules() {
		if module.IsEnabled() {
			if err := module.Generate(ctx, app.TemplateRoot(), app.OutputPath()); err != nil {
				fmt.Printf("[!] ERROR executing '%s': %s\n", module.Name(), err)
				continue
			}
			fmt.Printf("[+] module executed: %s\n", module.Name())
		} else {
			fmt.Printf("[-] disabled module: %s\n", module.Name())
		}
	}

	if err := viper.WriteConfig(); err != nil {
		fmt.Print(err)
	}
}

func registerModules() godin.ModuleRegistry {
	registry := &godin.DefaultRegistry{}

	err := registry.Register(grpcServer.NewGrpcServerModule())
	if err != nil {
		fmt.Print(err)
		os.Exit(-1)
	}

	return registry
}
