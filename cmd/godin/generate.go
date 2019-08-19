package main

import (
	"os"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"gitub.com/go-godin/godin"
)

func Generate(c *cli.Context) error {
	wd, _ := os.Getwd()
	app := godin.NewGodin(wd, "examples")

	// prepare configurator
	cfg := godin.NewConfigurator(app.RootPath())
	if !cfg.ConfigExists() {
		log.Fatal("not a godin project")
	}
	if err := cfg.Initialize(); err != nil {
		log.Fatal(err)
	}

	// resolve all enabled modules and configure them based on the godin.yaml
	if err := app.ResolveEnabledModules(godin.ModuleResolver{}, cfg); err != nil {
		log.WithError(err).Fatal("failed to resolve enabled modules")
	}

	// add modules to the configuration
	for _, module := range app.EnabledModules() {
		if err := cfg.Register(module); err != nil {
			log.WithError(err).Error("failed to register config-provider to config store")
		}
	}

	// parse a protobuf, which should also be passed via the cli
	ctx, err := godin.Parse(protobufPath)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to parse protobuf"))
	}
	log.Debugf("parsed protobuf: %s", protobufPath)

	// ensure the base output path exists
	if err := app.EnsureOutputPath(); err != nil {
		log.Println(err)
	}

	// TODO: load configuration via new configurator
	// TODO: use the resolver to resolve all enabled modules and return a preconfigured instance

	/*
		if err := app.InstallModule(godin.TransportGrpcServer); err != nil {
			log.Error(err)
		}
	*/

	if len(app.EnabledModules()) == 0 {
		log.Info("no modules enabled")
	}

	// generate all enabled modules
	for _, module := range app.EnabledModules() {
		log.Printf("==> Executing module '%s'\n", module.Identifier())

		// Generate() all modules
		if err := module.Generate(ctx, app.TemplateRoot(), app.OutputPath()); err != nil {
			log.Printf("[!] ERROR executing '%s': %s\n", module.Identifier(), err)
			continue
		}
	}

	if err := cfg.Save(); err != nil {
		log.WithError(err).Error("failed to save config-file")
	}

	return nil
}
