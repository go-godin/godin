package main

import (
	"fmt"
	"os"

	"gitub.com/go-godin/godin/module"

	"github.com/spf13/viper"
	"gitub.com/go-godin/godin"
)

func main() {
	wd, _ := os.Getwd()

	availableModules := registerModules()
	app := godin.NewGodin(availableModules, wd, "examples")

	if err := app.EnsureConfigFile(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// pretend to enable modules, this should be done via the godin cli
	if err := app.InstallModule(module.TransportGrpcServer); err != nil {
		fmt.Printf("ERROR installing module %s: %s\n", "transport.grpc.server", err)
		os.Exit(1)
	}
	fmt.Printf("==> Installed module %s\n", "transport.grpc.server")

	if err := app.InstallModule(module.TransportGrpcServer); err != nil {
		fmt.Printf("ERROR installing module %s: %s\n", "transport.grpc.server", err)
		os.Exit(1)
	}
	fmt.Printf("==> Installed module %s\n", "transport.grpc.server")

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
	for _, module := range app.EnabledModules() {
		fmt.Printf("==> Executing module '%s' with identifier '%s'\n", module.ConfigurationKey(), module.ID())
		if err := module.Generate(ctx, app.TemplateRoot(), app.OutputPath()); err != nil {
			fmt.Printf("[!] ERROR executing '%s (%s)': %s\n", module.ConfigurationKey(), module.ID(), err)
			continue
		}
	}
	fmt.Println("==> DONE")

	if err := viper.WriteConfig(); err != nil {
		fmt.Print(err)
	}
}

func registerModules() (modules []module.Type) {
	modules = append(
		modules,
		module.TransportGrpcServer,
	)
	return modules
}
