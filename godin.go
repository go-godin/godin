package godin

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/sirupsen/logrus"
)

const TemplateFolder = "templates"
const DefaultOutputFolder = "."

type Godin struct {
	enabledModules       Registry
	rootPath             string
	outputPath           string
	configRegistry       ConfigRegistry
	ProjectConfiguration *ProjectConfiguration
}

type ProjectConfiguration struct {
	Module   string
	Protobuf struct {
		Module string
	}
	Service struct {
		Namespace string
		Name      string
	}
	Templates struct {
		SourceFolder string `yaml:"sourceFolder"`
		OutputFolder string `yaml:"outputFolder"`
	}
}

type ConfigRegistry interface {
	Register(provider ConfigProvider) error
	Get(key string) interface{}
	Unmarshal(key string, target interface{}) error
}

type ModuleResolver interface {
	ResolveAll(source Resolvable) (modules []Module, err error)
}

// NewGodin returns a new, preconfigured, instance of godin.
// If outputPath is empty, the DefaultOutputFolder is used.
func NewGodin(configRegistry ConfigRegistry, rootPath string, outputPath string) *Godin {
	if outputPath == "" {
		outputPath = DefaultOutputFolder
	}

	cfg := &ProjectConfiguration{
		Templates: struct {
			SourceFolder string `yaml:"sourceFolder"`
			OutputFolder string `yaml:"outputFolder"`
		}{
			SourceFolder: TemplateFolder,
			OutputFolder: outputPath,
		},
	}

	g := &Godin{
		enabledModules:       NewRegistry(),
		rootPath:             rootPath,
		outputPath:           outputPath,
		configRegistry:       configRegistry,
		ProjectConfiguration: cfg,
	}

	return g
}

func NewGodinFromConfig(configRegistry ConfigRegistry, rootPath string) (*Godin, error) {
	g := &Godin{
		enabledModules: NewRegistry(),
		configRegistry: configRegistry,
		rootPath:       rootPath,
	}

	cfg := &ProjectConfiguration{}
	if err := configRegistry.Unmarshal(g.Identifier(), cfg); err != nil {
		return nil, err
	}
	g.ProjectConfiguration = cfg

	return g, nil
}

func (g *Godin) Identifier() string {
	return "project"
}

func (g *Godin) Configuration() interface{} {
	return g.ProjectConfiguration
}

// InstallModule adds a new module to the current project.
// The module will be looked up in the registry. If it exists, the Add() method is called
// to tell the module to add itself to the configuration.
func (g *Godin) InstallModule(module Module) error {
	if module == nil {
		return fmt.Errorf("failed to create module")
	}
	if err := module.Install(); err != nil {
		return err
	}

	if err := g.configRegistry.Register(module); err != nil {
		return fmt.Errorf("unable to register module with the configRegistry: %s", err)
	}

	if err := g.enabledModules.Register(module); err != nil {
		return err
	}

	return nil
}

func (g *Godin) ResolveEnabledModules(resolver ModuleResolver, cfg ResolvableConfig) error {
	enabledModules, err := resolver.ResolveAll(cfg)
	if err != nil {
		return err
	}

	for _, module := range enabledModules {
		logrus.Debugf("found enabled module '%s'", module.Identifier())
		if err := module.Configure(cfg); err != nil {
			return fmt.Errorf("unable to configure module '%s': %s", module.Identifier(), err)
		}
		if err := g.enabledModules.Register(module); err != nil {
			return fmt.Errorf("unable to register enabled module '%s': %s", module.Identifier(), err)
		}
		logrus.Debugf("configured module '%s'", module.Identifier())
	}
	return nil
}

// TemplateRoot returns the absolute path to the templates folder by joining the project's rootPath with the 'TemplateFolder'
func (g *Godin) TemplateRoot() string {
	return filepath.Join(g.rootPath, TemplateFolder)
}

// EnsureOutputPath ensures that the configured outputPath exists.
func (g *Godin) EnsureOutputPath() error {
	if _, err := os.Stat(filepath.Join(g.rootPath, g.outputPath)); os.IsNotExist(err) {
		if err := os.MkdirAll(g.outputPath, 0755); err != nil {
			return errors.Wrap(err, "EnsureOutputPath")
		}
	}
	return nil
}

// OutputPath returns the absolute path to the output path where all generated files are placed in.
func (g *Godin) OutputPath() string {
	return filepath.Join(g.rootPath, g.ProjectConfiguration.Templates.OutputFolder)
}

func (g *Godin) EnabledModules() Store {
	return g.enabledModules.Modules()
}

func (g *Godin) RootPath() string {
	return g.rootPath
}
