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

	// prepare configurator
	cfg := godin.NewConfigurator(wd)
	if !cfg.ConfigExists() {
		log.Fatal("not a godin project")
	}
	if err := cfg.Initialize(); err != nil {
		log.Fatal(err)
	}

	app, err := godin.NewGodinFromConfig(cfg, wd)
	if err != nil {
		log.WithError(err).Fatal("unable to load project")
	}

	// resolve all enabled modules and configure them based on the godin.yaml
	if err := app.ResolveEnabledModules(godin.NewModuleResolver(app.ProjectConfiguration), cfg); err != nil {
		log.WithError(err).Fatal("failed to resolve enabled modules")
	}

	// add godin to the configurator
	if err := cfg.Register(app); err != nil {
		log.WithError(err).Fatal("unable to register godin configuration")
	}

	// add existing modules to the configuration
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
		log.Error(err)
	}

	if len(app.EnabledModules()) == 0 {
		log.Info("no modules enabled")
	}

	// generate all enabled modules
	for _, module := range app.EnabledModules() {
		logger := log.WithField("module", module.Identifier())
		logger.Info("executing module")
		if err := module.Generate(app.ProjectConfiguration, ctx, app.TemplateRoot(), app.OutputPath()); err != nil {
			logger.WithError(err).Error("module generation failed")
			continue
		}
		logger.Info("module executed successfully")
	}

	if err := cfg.Save(); err != nil {
		log.WithError(err).Error("failed to save config-file")
	}

	return nil
}
