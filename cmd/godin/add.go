package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gitub.com/go-godin/godin"
	"gitub.com/go-godin/godin/module"
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

	if !g.ConfigExists() {
		log.Fatal("project not initialized")
	}

	switch moduleName {
	case "transport.grpc.server":
		log.Info("installing grpc server module")
		m := module.Factory(module.TransportGrpcServer)
		viper.Set(m.ConfigurationKey(), []interface{}{m.Configuration()})
		break
	default:
		log.Errorf("module '%s' is unknown, sorry", moduleName)
	}

	if err := viper.WriteConfig(); err != nil {
		log.Errorf("failed to write configuration: %s", err)
	}
	return nil
}
