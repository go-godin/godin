package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"gitub.com/go-godin/godin"

	"github.com/urfave/cli"
)

func SyncTemplates(c *cli.Context) error {
	force := c.Bool("force")
	wd, _ := os.Getwd()

	if force {
		log.Debug("force mode enabled")
	}

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

	// add godin to the configurator
	if err := cfg.Register(app); err != nil {
		log.WithError(err).Fatal("unable to register godin configuration")
	}

	// resolve all enabled modules and configure them based on the godin.yaml
	resolver := godin.NewModuleResolver(app.ProjectConfiguration)
	if err := app.ResolveEnabledModules(resolver, cfg); err != nil {
		log.WithError(err).Fatal("failed to resolve enabled modules")
	}

	return nil
}
