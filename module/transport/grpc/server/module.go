package server

import (
	"gitub.com/go-godin/godin"
)

// GrpcServerModule provides all templates for the gRPC server transport layer
type GrpcServerModule struct {
	*godin.BaseModule
	grpcServerConfig
	ServerTemplate godin.ModuleTemplate
}

// grpcServerConfig defines the transport.grpc.server module configuration struct.
type grpcServerConfig struct {
	Port    int    `mapstructure:"port"`
	Address string `mapstructure:"address"`
}

// NewGrpcServerModule returns a new pre-initialized GrpcServerModule.
// The module returns reasonable defaults so that it could be used directly.
// Upon calling Load(), the configuration will be loaded which eventually overwrites the defaults.
func NewGrpcServerModule() godin.Module {
	return &GrpcServerModule{
		BaseModule: &godin.BaseModule{
			ModuleName: "grpc-server",
		},
		grpcServerConfig: grpcServerConfig{
			Port:    50051,
			Address: "0.0.0.0",
		},
		ServerTemplate: NewServerTemplate(),
	}
}

// TODO: that's a crappy solution, find a better one!
func (mod *GrpcServerModule) New() godin.Module {
	return NewGrpcServerModule()
}

func (mod *GrpcServerModule) Install() error {
	return nil
}

// Generate will render the modules templates
func (mod *GrpcServerModule) Generate(ctx *godin.Context, templateRootPath, outputRootPath string) error {
	if err := mod.ServerTemplate.Render(ctx, mod.grpcServerConfig, templateRootPath, outputRootPath); err != nil {
		return err
	}
	return nil
}

// Configuration returns the module configuration as interface.
func (mod *GrpcServerModule) Configuration() interface{} {
	return mod.grpcServerConfig
}

// ConfigurationKey returns the GrpcServerModule key. It's used as unique identifier and - though only internally - also
// used as configuration key. Note that the ConfigurationKey() must not always correspond to the configuration key used by the
// module.
func (cfg *grpcServerConfig) ConfigurationKey() string {
	return "transport.grpc.server"
}
