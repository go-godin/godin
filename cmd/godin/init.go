package main

import (
	"github.com/spf13/viper"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"gitub.com/go-godin/godin"
)

func Init(c *cli.Context) error {
	namespace := c.String("namespace")
	service := c.String("service")
	if namespace == "" {
		log.Fatal("namespace cannot be empty")
	}
	if service == "" {
		log.Fatal("service cannot be empty")
	}

	wd, _ := os.Getwd()
	g := godin.NewGodin(wd, "")

	if g.ConfigExists() {
		log.Fatal("project already initialized")
	}

	if err := g.EnsureConfigFile(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Debug("created config file")


	log.Debugf("namespace: %s", namespace)
	log.Debugf("service: %s", service)

	viper.Set("project.namespace", namespace)
	viper.Set("project.service", service)

	if err := viper.WriteConfig(); err != nil {
		log.Errorf("failed to write configuration: %s", err)
	}
	log.Info("project initialized")

	return nil
}
