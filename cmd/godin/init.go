package main

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"gitub.com/go-godin/godin"
)

func Init(c *cli.Context) error {

	namespace := c.String("namespace")
	service := c.String("service")
	module := c.String("module")
	protoModule := c.String("protobuf-module")

	wd, _ := os.Getwd()

	cfg := godin.NewConfigurator(wd)
	g := godin.NewGodin(cfg, wd, "")

	if cfg.ConfigExists() {
		log.Fatal("already initialized")
	}
	if err := cfg.EnsureConfigFile(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	if err := cfg.Initialize(); err != nil {
		log.Fatal(err)
	}

	// register project config
	if err := cfg.Register(g); err != nil {
		log.Fatal("failed to register project configuration")
	}

	output, _ := filepath.Rel(wd, g.OutputPath())
	source, _ := filepath.Rel(wd, g.TemplateRoot())
	g.ProjectConfiguration.Module = module
	g.ProjectConfiguration.Protobuf.Module = protoModule
	g.ProjectConfiguration.Service.Name = service
	g.ProjectConfiguration.Service.Namespace = namespace
	g.ProjectConfiguration.Templates.OutputFolder = output
	g.ProjectConfiguration.Templates.SourceFolder = source

	log.Debug("created config file")

	if err := cfg.Save(); err != nil {
		log.Errorf("failed to write configuration: %s", err)
	}
	log.Info("project initialized")

	return nil
}
