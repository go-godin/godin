package godin

import (
	"fmt"
	"os"
	"path"

	"github.com/pkg/errors"

	"github.com/spf13/viper"
)

type ConfigProvider interface {
	// Identifier returns the unique identifier by which a module can be distinguished from others.
	// An identifier is a '.' separated namespace which will be resolved to the config hierarchy.
	Identifier() string

	// Configuration returns the configuration of the ConfigProvider object.
	Configuration() interface{}
}

const ConfigFileName = "godin"
const ConfigFileType = "yaml"

type Configurator struct {
	configRegistry map[string]ConfigProvider
	RootPath       string
}

func NewConfigurator(configPath string) *Configurator {
	return &Configurator{
		configRegistry: make(map[string]ConfigProvider),
		RootPath:       configPath,
	}
}

func (cfg *Configurator) Initialize() error {
	viper.AddConfigPath(cfg.RootPath)
	viper.SetConfigName(ConfigFileName)
	viper.SetConfigType(ConfigFileType)
	return viper.ReadInConfig()
}

func (cfg *Configurator) IsSet(key string) bool {
	return viper.IsSet(key)
}

func (cfg *Configurator) Get(key string) interface{} {
	return viper.Get(key)
}

func (cfg *Configurator) Unmarshal(key string, target interface{}) error {
	return viper.UnmarshalKey(key, target)
}

func (cfg *Configurator) Register(provider ConfigProvider) error {
	if existing := cfg.configRegistry[provider.Identifier()]; existing != nil {
		return fmt.Errorf("a config provider with identifier '%s' is already registered", provider.Identifier())
	}
	cfg.configRegistry[provider.Identifier()] = provider

	return nil
}

func (cfg *Configurator) Save() error {
	for _, reg := range cfg.configRegistry {
		viper.Set(reg.Identifier(), reg.Configuration())
	}
	return viper.WriteConfig()
}

// ConfigExists checks whether a configuration file exists. That's the indicator whether a project
// has been initialized.
func (cfg *Configurator) ConfigExists() bool {
	p := path.Join(cfg.RootPath, fmt.Sprintf("%s.%s", ConfigFileName, ConfigFileType))

	if _, err := os.Stat(p); err != nil {
		return false
	}
	return true
}

// EnsureConfigFile a godin project directory in the configured 'rootPath' and ensure a configuration file exists.
// If the project is already initialized, nothing is returned (silent fail) which makes this method idempotent.
// Note: If a config file is created, it will be empty.
func (cfg *Configurator) EnsureConfigFile() error {
	if cfg.ConfigExists() {
		return nil
	}
	_, err := os.Create(path.Join(cfg.RootPath, fmt.Sprintf("%s.%s", ConfigFileName, ConfigFileType)))
	if err != nil {
		return errors.Wrap(err, "failed to initialize project")
	}

	return nil
}
