package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"gitub.com/go-godin/godin"

	"github.com/urfave/cli"
)

func Add(c *cli.Context) error {
	moduleName := c.Args().Get(0)
	if moduleName == "" {
		log.Fatal("the module name must be passed")
	}
	if force {
		log.Debug("force mode enabled")
	}

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
	resolver := godin.NewModuleResolver(app.ProjectConfiguration)
	if err := app.ResolveEnabledModules(resolver, cfg); err != nil {
		log.WithError(err).Fatal("failed to resolve enabled modules")
	}

	// add godin to the configurator
	if err := cfg.Register(app); err != nil {
		log.WithError(err).Fatal("unable to register godin configuration")
	}

	// add existing modules to the configurator
	for _, module := range app.EnabledModules() {
		if err := cfg.Register(module); err != nil {
			log.WithError(err).Error("unable to register config-provider to config store")
		}
	}

	newModule, err := resolver.Resolve(moduleName)
	if err != nil {
		log.WithError(err).Fatal("unable to resolve module name")
	}

	if err := app.InstallModule(newModule); err != nil {
		log.WithError(err).Fatal("unable to install module")
	}

	// write module templates into project templates folder
	writer := godin.NewTemplateWriter(godin.Templates)
	for _, tpl := range newModule.Templates() {
		tplSource := tpl.Configuration().SourceFile
		logger := log.WithFields(log.Fields{
			"template": tplSource,
			"module":   newModule.Identifier(),
		})

		if force {
			logger.Warning("overwriting existing template")
			copyTemplate(logger, app.TemplateRoot(), tplSource, writer, true)
		} else {
			copyTemplate(logger, app.TemplateRoot(), tplSource, writer, false)
		}
	}

	if err := cfg.Save(); err != nil {
		log.WithError(err).Error("unable to save config-file")
	}

	log.WithField("module", newModule.Identifier()).Info("module installed")

	return nil
}
