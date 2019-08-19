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

	wd, _ := os.Getwd()

	// prepare configurator
	cfg := godin.NewConfigurator(wd)
	if !cfg.ConfigExists() {
		log.Fatal("not a godin project")
	}
	if err := cfg.Initialize(); err != nil {
		log.Fatal(err)
	}

	app := godin.NewGodin(cfg, wd, "")

	// resolve all enabled modules and configure them based on the godin.yaml
	resolver := godin.ModuleResolver{}
	if err := app.ResolveEnabledModules(resolver, cfg); err != nil {
		log.WithError(err).Fatal("failed to resolve enabled modules")
	}

	// add existing modules to the configuration
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

	if err := cfg.Save(); err != nil {
		log.WithError(err).Error("unable to save config-file")
	}

	/*

		// TODO: ensure the module isn't already installed

		switch moduleName {
		case "transport.grpc.server":
			log.Info("installing grpc server module")
			m := godin.ModuleFactory(godin.TransportGrpcServer)
			viper.Set(m.Identifier(), []interface{}{m.Configuration()})
			break
		default:
			log.Errorf("module '%s' is unknown, sorry", moduleName)
		}

		if err := viper.WriteConfig(); err != nil {
			log.Errorf("failed to write configuration: %s", err)
		}

	*/
	return nil
}
