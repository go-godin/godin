package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gitub.com/go-godin/godin"
	"os"

	"github.com/urfave/cli"
)

func Add(c *cli.Context) error {
	moduleName := c.Args().Get(0)
	if moduleName == "" {
		log.Fatal("the module name must be passed")
	}

	wd, _ := os.Getwd()
	g := godin.NewGodin(wd, "")

	// prepare configurator
	cfg := godin.Configurator{RootPath: g.RootPath()}

	if !cfg.ConfigExists() {
		log.Fatal("not a godin project")
	}
	if err := cfg.Initialize(); err != nil {
		log.Fatal(err)
	}

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
	return nil
}
