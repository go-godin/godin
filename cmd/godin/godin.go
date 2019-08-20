package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli"
)

var protobufPath string
var debug bool

func main() {

	app := cli.NewApp()
	app.Name = "Godin"
	app.Usage = "A go-kit based microservice toolkit"
	app.Version = "1.0.0"
	app.EnableBashCompletion = true
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "debug, d",
			Usage:       "include debug logs",
			EnvVar:      "DEBUG",
			Required:    false,
			Destination: &debug,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "namespace, n",
					Usage:    "The namespace of the project (e.g. 'user' for the user stack)",
					Required: true,
				},
				cli.StringFlag{
					Name:     "service, s",
					Usage:    "The service name which this project owns, normalized (lowercase)",
					Required: true,
				},
			},
			Usage:  "Initialize a godin project in the current directory",
			Action: Init,
			Before: setLogLevel,
		},
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "Add a module to the current project",
			Before:  setLogLevel,
			Action:  Add,
		},
		{
			Name:    "generate",
			Aliases: []string{"g"},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "protobuf, p",
					Usage:       "Absolute path to the target protobuf",
					EnvVar:      "PROTOBUF_FILE",
					Required:    true,
					Destination: &protobufPath,
				},
			},
			Usage:  "generate all enabled module templates",
			Action: Generate,
			Before: setLogLevel,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func setLogLevel(c *cli.Context) error {
	if debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
	return nil
}
