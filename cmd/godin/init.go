package main

import (
	"os"

	"github.com/spf13/viper"

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

	cfg := godin.NewConfigurator(wd)
	//g := godin.NewGodin(cfg, wd, "")

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
