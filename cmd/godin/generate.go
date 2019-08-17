package main

import (
	"os"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
	"gitub.com/go-godin/godin"
)

func Generate() func(c *cli.Context) error {
	wd, _ := os.Getwd()
	app := godin.NewGodin(wd, "examples")

	if err := app.EnsureConfigFile(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Debug("configuration found, godin project initialized")

	// parse a protobuf, which should also be passed via the cli
	ctx, err := godin.Parse(protobufPath)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to parse protobuf"))
	}
	log.Debugf("parsed protobuf %s", protobufPath)

	// ensure we can write the generated files
	if err := app.EnsureOutputPath(); err != nil {
		log.Println(err)
	}

	// generate all enabled modules
	if len(app.EnabledModules()) == 0 {
		log.Info("no modules enabled")
	}
	for _, module := range app.EnabledModules() {
		log.Printf("==> Executing module '%s' with identifier '%s'\n", module.ConfigurationKey(), module.ID())
		if err := module.Generate(ctx, app.TemplateRoot(), app.OutputPath()); err != nil {
			log.Printf("[!] ERROR executing '%s (%s)': %s\n", module.ConfigurationKey(), module.ID(), err)
			continue
		}
	}

	if err := viper.WriteConfig(); err != nil {
		log.Print(err)
	}

	return nil
}
