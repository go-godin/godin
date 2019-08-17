package godin

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// Configurable defines the interface of anything which can be configured must behave.
type Configurable interface {
	Key() string
	Load() error
	Save()
	Delete()
}

const ConfigFile = "godin"
const ConfigFileType = "yaml"
const TemplateFolder = "templates"
const DefaultOutputFolder = "."

type Godin struct {
	modules    ModuleRegistry
	rootPath   string
	outputPath string
}

// NewGodin returns a new, preconfigured, instance of godin.
func NewGodin(registry ModuleRegistry, rootPath string, outputPath string) *Godin {
	viper.AddConfigPath(rootPath)
	viper.SetConfigName(ConfigFile)
	viper.SetConfigType(ConfigFileType)

	if outputPath == "" {
		outputPath = DefaultOutputFolder
	}

	g := &Godin{
		modules:    registry,
		rootPath:   rootPath,
		outputPath: outputPath,
	}

	return g
}

// AddModule adds a new module to the current project.
// The module will be looked up in the registry. If it exists, the Add() method is called
// to tell the module to add itself to the configuration.
func (g *Godin) AddModule(name string) error {
	module, err := g.modules.Get(name)
	if err != nil {
	    return errors.Wrap(err, "failed to add module")
	}
	module.Add()

	return nil
}

// ConfigExists checks whether a configuration file exists. That's the indicator whether a project
// has been initialized.
func (g *Godin) ConfigExists() bool {
	p := path.Join(g.rootPath, fmt.Sprintf("%s.%s", ConfigFile, ConfigFileType))
	fmt.Println(p)

	if _, err := os.Stat(p); err != nil {
		return false
	}
	return true
}

// TemplateRoot returns the absolute path to the templates folder by joining the project's rootPath with the 'TemplateFolder'
func (g *Godin) TemplateRoot() string {
	return filepath.Join(g.rootPath, TemplateFolder)
}

// EnsureOutputPath ensures that the configured outputPath exists.
func (g *Godin) EnsureOutputPath() error {
	if _, err := os.Stat(filepath.Join(g.rootPath, g.outputPath)); os.IsNotExist(err) {
		if err := os.MkdirAll(g.outputPath, 0755); err != nil {
			return err
		}
	}
	return nil
}

// OutputPath returns the absolute path to the output path where all generated files are placed in.
func (g *Godin) OutputPath() string {
	return filepath.Join(g.rootPath, g.outputPath)
}

// EnsureConfigFile a godin project directory in the configured 'rootPath' and ensure a configuration file exists.
// If the project is already initialized, nothing is returned (silent fail) which makes this method idempotent.
// Note: If a config file is created, it will be empty.
func (g *Godin) EnsureConfigFile() error {
	if g.ConfigExists() {
		return nil
	}
	_, err := os.Create(path.Join(g.rootPath, fmt.Sprintf("%s.%s", ConfigFile, ConfigFileType)))
	if err != nil {
		return errors.Wrap(err, "failed to initialize project")
	}

	return nil
}
