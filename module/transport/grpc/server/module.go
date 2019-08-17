package server

import (
	"fmt"

	"github.com/spf13/viper"
	"gitub.com/go-godin/godin"
)

// GrpcServerModule provides all templates for the gRPC server transport layer
type GrpcServerModule struct {
	GrpcServerConfiguration
	ServerTemplate godin.ModuleTemplate
}

// GrpcServerConfiguration defines the transport.grpc.server module configuration struct.
type GrpcServerConfiguration struct {
	Port    int    `mapstructure:"port"`
	Address string `mapstructure:"address"`
}

var grpcServerTemplates = []*godin.TemplateConfiguration{
	{
		Name:       "grpc-server",
		SourceFile: "transport/grpc/server.go.tmpl",
		TargetFile: godin.DefaultTargetFile("transport/grpc/server.go.tmpl"),
		GoSource:   true,
		Skip:       false,
	},
	{
		Name:       "grpc-server-1",
		SourceFile: "transport/grpc/server1.go.tmpl",
		TargetFile: godin.DefaultTargetFile("transport/grpc/server.go.tmpl"),
		GoSource:   true,
		Skip:       false,
	},
	{
		Name:       "grpc-server-2",
		SourceFile: "transport/grpc/server1.go.tmpl",
		TargetFile: godin.DefaultTargetFile("transport/grpc/server.go.tmpl"),
		GoSource:   true,
		Skip:       true,
	},
}

// NewGrpcServerModule returns a new pre-initialized GrpcServerModule.
// The module returns reasonable defaults so that it could be used directly.
// Upon calling Load(), the configuration will be loaded which eventually overwrites the defaults.
func NewGrpcServerModule() godin.Module {
	return &GrpcServerModule{
		GrpcServerConfiguration: GrpcServerConfiguration{
			Port:    50051,
			Address: "0.0.0.0",
		},
		ServerTemplate: NewServerTemplate(),
	}
}

func (mod *GrpcServerModule) Generate(ctx *godin.Context, templateRootPath, outputRootPath string) error {
	if err := mod.ServerTemplate.Render(ctx, mod.GrpcServerConfiguration, templateRootPath, outputRootPath); err != nil {
		return err
	}
	return nil
}

func (mod *GrpcServerModule) PrepareContext(ctx *godin.Context) interface{} {
	return struct {
		CTX        *godin.Context
		GrpcServer *GrpcServerConfiguration
	}{
		CTX:        ctx,
		GrpcServer: &mod.GrpcServerConfiguration,
	}
}

// Name returns a printable name of the module.
func (mod *GrpcServerModule) Name() string {
	return "godin:transport:grpc:server"
}

// Addable will set the module into the enabled state. In this case, enabled means that the configuration exists.
func (mod *GrpcServerModule) Add() {
	mod.Save()
}

// Removable will disable the module by removing it from the configuration
func (mod *GrpcServerModule) Remove() {
	if mod.IsEnabled() {
		mod.Delete()
	}
}

// IsEnabled returns true if the module is enabled. In this case, if the config key is set, then
// the module is considered enabled.
func (mod *GrpcServerModule) IsEnabled() bool {
	return viper.IsSet(mod.Key())
}

// Key returns the GrpcServerModule key. It's used as unique identifier and - though only internally - also
// used as configuration key. Note that the Key() must not always correspond to the configuration key used by the
// module.
func (cfg *GrpcServerConfiguration) Key() string {
	return "transport.grpc.server"
}

// Load will load the transport.grpc.server sub-tree from the config file.
// The data is then directly marshalled into the current GrpcServerConfiguration instance on the fly.
func (cfg *GrpcServerConfiguration) Load() error {
	if err := viper.UnmarshalKey(cfg.Key(), cfg); err != nil {
		return fmt.Errorf("failed to unmarshal GrpcServerConfiguration '%s': %s", cfg.Key(), err)
	}

	return nil
}

func (cfg *GrpcServerConfiguration) Save() {
	viper.Set(cfg.Key(), cfg)
}

func (cfg *GrpcServerConfiguration) Delete() {
	delete(viper.Get(cfg.Key()).(map[string]interface{}), "server")
	cfg.Save()
}
