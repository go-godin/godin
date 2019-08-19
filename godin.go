package godin

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

const TemplateFolder = "templates"
const DefaultOutputFolder = "."

type Godin struct {
	enabledModules Registry
	rootPath       string
	outputPath     string
	configurator   *Configurator
}

// NewGodin returns a new, preconfigured, instance of godin.
// If outputPath is empty, the DefaultOutputFolder is used.
func NewGodin(configurator *Configurator, rootPath string, outputPath string) *Godin {
	if outputPath == "" {
		outputPath = DefaultOutputFolder
	}

	g := &Godin{
		enabledModules: NewRegistry(),
		rootPath:       rootPath,
		outputPath:     outputPath,
		configurator:   configurator,
	}

	return g
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

	if err := g.configurator.Register(module); err != nil {
		return fmt.Errorf("unable to register module with the configurator: %s", err)
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
			return err
		}
	}
	return nil
}

// OutputPath returns the absolute path to the output path where all generated files are placed in.
func (g *Godin) OutputPath() string {
	return filepath.Join(g.rootPath, g.outputPath)
}

func (g *Godin) EnabledModules() Store {
	return g.enabledModules.Modules()
}

func (g *Godin) RootPath() string {
	return g.rootPath
}
