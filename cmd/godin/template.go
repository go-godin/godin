package main

import (
	"os"
	"path"

	log "github.com/sirupsen/logrus"
	"gitub.com/go-godin/godin"

	"github.com/urfave/cli"
)

func SyncTemplates(c *cli.Context) error {
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

	writer := godin.NewTemplateWriter(godin.Templates)
	for _, module := range app.EnabledModules() {
		for _, template := range module.Templates() {
			tplSource := template.Configuration().SourceFile
			logger := log.WithFields(log.Fields{
				"template": tplSource,
				"module":   module.Identifier(),
			})
			// template exists
			if template.Configuration().SourceExists(app.TemplateRoot()) {
				if force {
					logger.Warning("overwriting existing template")
					copyTemplate(logger, app.TemplateRoot(), tplSource, writer, true)
				} else {
					logger.Debug("skipping existing template")
				}
				continue
			}
			copyTemplate(logger, app.TemplateRoot(), tplSource, writer, false)
		}
	}
	return nil
}

func copyTemplate(logger *log.Entry, templateRoot, templateSource string, writer *godin.TemplateWriter, overwrite bool) {
	tplSourceAbs := path.Join(templateRoot, templateSource)
	if err := writer.EnsurePath(tplSourceAbs); err != nil {
		logger.WithError(err).Error("unable to ensure template path")
		return
	}

	if overwrite {
		if err := writer.OverWrite(templateSource, tplSourceAbs); err != nil {
			logger.WithError(err).Warning("unable to write template, module may not work correctly")
			return
		}
	} else {
		if err := writer.Write(templateSource, tplSourceAbs); err != nil {
			logger.WithError(err).Warning("unable to write template, module may not work correctly")
			return
		}
	}
	logger.Info("template written")
}
