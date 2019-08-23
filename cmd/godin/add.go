package main

import (
	"os"
	"path"

	log "github.com/sirupsen/logrus"
	"gitub.com/go-godin/godin"

	"github.com/urfave/cli"
)

func Add(c *cli.Context) error {
	moduleName := c.Args().Get(0)
	if moduleName == "" {
		log.Fatal("the module name must be passed")
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
		tplTargetAbs := path.Join(app.TemplateRoot(), tplSource)
		logger := log.WithFields(log.Fields{
			"sourceFile": "binary://" + tplSource,
			"targetFile": tplSource,
		})

		if err := writer.EnsurePath(tplTargetAbs); err != nil {
			logger.WithError(err).Error("unable to ensure template path")
			continue
		}
		logger.Debug("template target exists")

		if err := writer.Write(tplSource, tplTargetAbs); err != nil {
			logger.WithError(err).Error("unable to write template, module may not work correctly")
			continue
		}
		logger.Info("template written")
	}

	if err := cfg.Save(); err != nil {
		log.WithError(err).Error("unable to save config-file")
	}

	return nil
}
